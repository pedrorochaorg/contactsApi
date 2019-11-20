package contactsApi

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/pedrorochaorg/contactsApi/db"
	"github.com/pedrorochaorg/contactsApi/obj"
	"github.com/DATA-DOG/go-sqlmock"

)


