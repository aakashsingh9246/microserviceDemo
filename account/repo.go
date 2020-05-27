package account

import(
	"github.com/gocql/gocql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type repo struct {
	db     *gocql.Session
	logger log.Logger
}

func NewRepo(db *gocql.Session, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "cassandra"),
	}
}

func (repo *repo) CreateUser(user User) error {
	fmt.Println(" **** Creating new emp ****\n", emp)
	if err := repo.db.Query("INSERT INTO emp(emp_id, emp_city, emp_name, emp_phone, emp_sal) VALUES(?, ?, ?, ?, ?)",
		user.ID, user.City, user.Name, user.Phone, user.Sal).Exec(); err != nil {
		fmt.Println("Error while inserting Emp")
		return err
	}
	return nil
}

func (repo *repo) GetUser(id string) (string, error) {
	fmt.Println("Getting Employees")

	m := map[string]interface{}{}
	var user User
	iter := repo.db.Query("SELECT * FROM emp emp_id = ?",id).Iter()
	for iter.MapScan(m) {
		user = User{
			m["emp_id"],
			m["emp_name"],
			m["emp_city"],
			m["emp_phone"],
			m["emp_sal"],
		}
	}

	return 	user.Name, nil
}