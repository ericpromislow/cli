package helpers

import (
	"bufio"
	"bytes"
	"io"
	"net/http"
	"strings"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

func VerifyRequestFromDumpSttring(input string) http.HandlerFunc {
	buf := bufio.NewReader(strings.NewReader(input))
	req, err := http.ReadRequest(buf)
	Expect(err).ToNot(HaveOccurred())
	return ghttp.CombineHandlers(
		ghttp.VerifyRequest(req.Method, req.URL.Path, req.URL.RawQuery),
		ghttp.VerifyHeader(req.Header),
		ghttp.VerifyBody([]byte(bodyToString(req.Body))),
	)
}

func RespondWithDumpString(input string) http.HandlerFunc {
	buf := bufio.NewReader(strings.NewReader(input))
	resp, err := http.ReadResponse(buf, nil)
	Expect(err).ToNot(HaveOccurred())
	return ghttp.RespondWith(resp.StatusCode, bodyToString(resp.Body), resp.Header)
}

func bodyToString(i io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(i)
	return buf.String()
}
