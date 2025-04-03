import (
	"context"
	"errors"
	"fmt"
	{{if .time}}"time"{{end}}
	{{if .containsDbSql}}"database/sql"{{end}}
	"github.com/tao-kit/gormc"
    "github.com/tao-kit/gormc/batchx"
	"github.com/tao-kit/gormc/pagex"
	"github.com/tao-kit/tao/core/stores/cache"
	"gorm.io/gorm"

	{{if .third}}{{.third}}{{end}}
)
