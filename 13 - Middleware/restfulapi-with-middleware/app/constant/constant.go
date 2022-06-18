package constant

import "encoding/base64"

var SECRET_JWT string = base64.StdEncoding.EncodeToString([]byte("golang"))
