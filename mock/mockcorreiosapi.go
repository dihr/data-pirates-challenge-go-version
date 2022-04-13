package mock

import (
	"data-pirates-challenge-go-version/services/correiosapi"
	"io"
	"net/http"
	"strings"
)

type mockCorreriosApi struct {
}

func NewMockCorreiosApi() correiosapi.CorreiosApiSvc {
	return &mockCorreriosApi{}
}

func (m mockCorreriosApi) GetHomePage() (*http.Response, error) {
	htmlBody := `<select name=UF class="f1col">
<option value=""</option>
<option value="AC">AC</option>
<option value="AL">AL</option>
<option value="AM">AM</option>
<option value="AP">AP</option>
</select>`
	return &http.Response{Body: io.NopCloser(strings.NewReader(htmlBody))}, nil
}

func (m mockCorreriosApi) GetData(uf string, startPage interface{}, endPage interface{}) (*http.Response, error) {
	htmlBody := `
<table class="tmptabela">
<tr bgcolor="#C4DEE9">
<td width="100">Acrelândia</td>
<td width="80"> 69945-000 a 69949-999</td>	
<td width="100">N&atilde;o codificada por logradouros</td>	
<td width="85">Total do munic&iacute;pio</td>
</tr>
</table>`
	return &http.Response{Body: io.NopCloser(strings.NewReader(htmlBody))}, nil
}
