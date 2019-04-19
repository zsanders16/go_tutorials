package form_test

import (
	"html/template"
	"testing"

	"github.com/zsanders16/go_tutorials/testing/form"
)

var (
	tplTypeNameValue = template.Must(template.New("").Parse(`<input type="{{.Type}}" name="{{.Name}}"{{with .Value}} value="{{.}}"{{end}}>`))
)

func TestHTML(t *testing.T) {
	tests := map[string]struct {
		tpl     *template.Template
		strct   interface{}
		want    template.HTML
		wantErr error
	}{
		"A basic form with values": {
			tpl: tplTypeNameValue,
			strct: struct {
				Name  string
				Email string
			}{
				Name:  "Zack Sanders",
				Email: "zack@test.com",
			},
			want: `<input type="text" name="Name" value="Zack Sanders">` +
				`<input type="text" name="Email" value="zack@test.com">`,
			wantErr: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, gotErr := form.HTML(tc.tpl, tc.strct)
			if gotErr != tc.wantErr {
				t.Fatalf("HTML() err = %v; want %v", gotErr, tc.wantErr)
			}

			if got != tc.want {
				t.Errorf("HTML() = %q; want %q", got, tc.want)
			}
		})
	}
}
