<div align="center" style="padding: 20px 0;">
  <img src=".github/rest.svg" alt="Rest Logo" />
</div>

<div align="center">
  <img alt="GitHub Release" src="https://img.shields.io/github/v/release/ogabrielrodrigues/rest?display_name=release">
  <img alt="GitHub License" src="https://img.shields.io/github/license/ogabrielrodrigues/rest">
</div>

<hr>

<h3 style="font-size: 28px; text-decoration:none;">Tired of having to do everything by hand every time you start a new go project with the <a href="https://github.com/go-chi/chi" target="_blank">chi</a> router?</h3>

<p style="font-size: 20px; text-decoration:none;">This set of pre-ready functions will help you save time and make you more productive.</p>
<br>
<p style="font-size: 16px; text-decoration:none;">
  <code>Binding</code>
  <br> 
  Request body bind and validation with 
  <a href="https://github.com/go-playground/validator" target="_blank">Validator</a>.
</p>

```go
import "github.com/ogabrielrodrigues/rest"

type Body struct {
  Name string `json:"name" validate:"required,min=2"`
  Age int `json:"age" validate:"required,min=1,max=140"`
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
  body := Body{}

  if err := rest.Bind(r.Body, &body); err != nil {
    rest.JSON(w, err.Code, err)
    return
  }

  // Rest of your handler logic...
}
```

<br>
<hr>
<br>
<p style="font-size: 16px; text-decoration:none;">
  <code>Error handling</code> 
  <br>
  Better handling of major http status errors.
</p>

```go
import "github.com/ogabrielrodrigues/rest"

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
  rest_error := rest.NewInternalServerErr(
    "this error occurred because the logic is not ready",
  )

  rest.JSON(w, rest_error.Code, rest_error)
}

// Output in JSON response:
{
  "message": "this error occurred because the logic is not ready",
  "code": 500,
  "error": "internal_server_error"
}
```

<br>
<hr>
<br>
<p style="font-size: 16px; text-decoration:none;">
  <code>Response</code> 
  <br>
  Simple way to return data with <code>ResponseWriter</code>.
</p>

```go
import (
  "net/http"
  "github.com/ogabrielrodrigues/rest"
)

func ResponseJSONHandler(w http.ResponseWriter, r *http.Request) {
  response := map[string]string{
    "message": "Thank you if you are reading this docs!!",
  }

  rest.JSON(w, http.StatusOK, response)
}

// Output in JSON response:
{
  "message": "Thank you if you are reading this docs!!",
}
```

<p style="font-size: 16px; text-decoration:none;">
  Or prefer not to return data:
</p>

```go
import (
  "net/http"
  "github.com/ogabrielrodrigues/rest"
)

func OnlyStatusHandler(w http.ResponseWriter, r *http.Request) {
  rest.End(w, http.StatusOK)
}

// Output in response headers:
HTTP/1.1 200 OK
```

<br>
<hr>
<br>
<p style="font-size: 16px; text-decoration:none;">
  <code>Validation</code> 
  <br>
  Validate your request body or URL params with
  <a href="https://github.com/go-playground/validator" target="_blank">Validator</a>.
</p>

```go
import "github.com/ogabrielrodrigues/rest"

type Body struct {
  Name string `json:"name" validate:"required,min=2"`
  Age int `json:"age" validate:"required,min=1,max=140"`
}

func StructValidationHandler(w http.ResponseWriter, r *http.Request) {
  body := Body{}

  if err := rest.Validate.Struct(&body); err != nil {
    rest_error := rest.ValidateStructErr(err)
    rest.JSON(w, rest_error.Code, rest_error)
    return
  }

  // Rest of your handler logic...
}
```

<p style="font-size: 16px; text-decoration:none;">
  Or validate one URL param:
</p>

```go
import (
  "github.com/go-chi/chi/v5"
  "github.com/ogabrielrodrigues/rest"
)

func VarValidationHandler(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

  if err := rest.Validate.Var(id, "uuid4"); err != nil {
    rest_error := rest.ValidateVarErr(err)
    rest.JSON(w, rest_error.Code, rest_error)
    return
  }

  // Rest of your handler logic...
}
```

<br>
<h3 align="center" style="font-size: 20px; text-decoration:none;">Make with ❤️ by <a href="https://github.com/ogabrielrodrigues">ogabrielrodrigues</a></h3>
