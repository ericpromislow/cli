package v7action

import (
	"errors"
	"strings"
	"time"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"

	log "github.com/sirupsen/logrus"
)

type Build struct {
	GUID string
}

func (actor Actor) StagePackage(packageGUID, appName, spaceGUID string) (<-chan Droplet, <-chan Warnings, <-chan error) {
	dropletStream := make(chan Droplet)
	warningsStream := make(chan Warnings)
	errorStream := make(chan error)
	go func() {
		defer close(dropletStream)
		defer close(warningsStream)
		defer close(errorStream)

		apps, warnings, err := actor.GetApplicationsByNamesAndSpace([]string{appName}, spaceGUID)
		warningsStream <- warnings
		if err != nil {
			if _, ok := err.(actionerror.ApplicationsNotFoundError); ok {
				err = actionerror.ApplicationNotFoundError{Name: appName}
			}
			errorStream <- err
			return
		}
		app := apps[0]

		pkgs, allWarnings, err := actor.CloudControllerClient.GetPackages(ccv3.Query{
			Key: ccv3.AppGUIDFilter, Values: []string{app.GUID},
		})
		warningsStream <- Warnings(allWarnings)
		if err != nil {
			errorStream <- err
			return
		}

		if len(pkgs) == 0 {
			err = actionerror.PackageNotFoundInAppError{GUID: packageGUID, AppName: appName}
			errorStream <- err
			return
		}

		build := ccv3.Build{PackageGUID: packageGUID}
		build, allWarnings, err = actor.CloudControllerClient.CreateBuild(build)
		warningsStream <- Warnings(allWarnings)

		if err != nil {
			errorStream <- err
			return
		}

		timer := actor.Clock.NewTimer(time.Millisecond)
		defer timer.Stop()
		timeout := actor.Clock.After(actor.Config.StagingTimeout())

		for {
			select {
			case <-timeout:
				errorStream <- actionerror.StagingTimeoutError{AppName: appName, Timeout: actor.Config.StagingTimeout()}
				return
			case <-timer.C():
				var warnings ccv3.Warnings
				build, warnings, err = actor.CloudControllerClient.GetBuild(build.GUID)
				warningsStream <- Warnings(warnings)
				if err != nil {
					errorStream <- err
					return
				}

				switch build.State {
				case constant.BuildFailed:
					if strings.Contains(build.Error, "NoAppDetectedError") {
						errorStream <- actionerror.StagingFailedNoAppDetectedError{Reason: build.Error}
					} else {
						errorStream <- actionerror.StagingFailedError{Reason: build.Error}
					}
					return
				case constant.BuildStaging:
					timer.Reset(actor.Config.PollingInterval())
				default:

					//TODO: uncomment after #150569020
					// ccv3Droplet, warnings, err := actor.CloudControllerClient.GetDroplet(build.DropletGUID)
					// warningsStream <- Warnings(warnings)
					// if err != nil {
					// 	errorStream <- err
					// 	return
					// }

					ccv3Droplet := ccv3.Droplet{
						GUID:      build.DropletGUID,
						State:     constant.DropletState(build.State),
						CreatedAt: build.CreatedAt,
					}

					dropletStream <- actor.convertCCToActorDroplet(ccv3Droplet)
					return
				}
			}
		}
	}()

	return dropletStream, warningsStream, errorStream
}

func (actor Actor) StageApplicationPackage(packageGUID string) (Build, Warnings, error) {
	var allWarnings Warnings

	build := ccv3.Build{PackageGUID: packageGUID}
	build, warnings, err := actor.CloudControllerClient.CreateBuild(build)
	log.Debug("created build")
	allWarnings = append(allWarnings, warnings...)
	if err != nil {
		return Build{}, allWarnings, err
	}

	log.Debug("no errors creating build")
	return Build{GUID: build.GUID}, allWarnings, nil
}

func (actor Actor) PollBuild(buildGUID string, appName string) (Droplet, Warnings, error) {
	var allWarnings Warnings

	timeout := actor.Clock.After(actor.Config.StagingTimeout())
	interval := actor.Clock.NewTimer(time.Millisecond)

	for {
		select {
		case <-interval.C():
			build, warnings, err := actor.CloudControllerClient.GetBuild(buildGUID)
			allWarnings = append(allWarnings, warnings...)
			if err != nil {
				return Droplet{}, allWarnings, err
			}

			switch build.State {
			case constant.BuildFailed:
				return Droplet{}, allWarnings, errors.New(build.Error)

			case constant.BuildStaged:
				droplet, warnings, err := actor.CloudControllerClient.GetDroplet(build.DropletGUID)
				allWarnings = append(allWarnings, warnings...)
				if err != nil {
					return Droplet{}, allWarnings, err
				}

				return Droplet{
					GUID:      droplet.GUID,
					State:     droplet.State,
					CreatedAt: droplet.CreatedAt,
				}, allWarnings, nil
			}

			interval.Reset(actor.Config.PollingInterval())

		case <-timeout:
			return Droplet{}, allWarnings, actionerror.StagingTimeoutError{AppName: appName, Timeout: actor.Config.StagingTimeout()}
		}
	}
}
