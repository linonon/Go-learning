package testify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

type User struct {
	Name string
	Age  int
}

type ICrawler interface {
	GetUserList() ([]*User, error)
}

type MyCrawler struct {
	url string
}

func (c *MyCrawler) GetUserList() ([]*User, error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userList []*User
	err = json.Unmarshal(data, &userList)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func GetAndPrintUsers(crawler ICrawler) {
	users, err := crawler.GetUserList()
	if err != nil {
		return
	}

	for _, u := range users {
		fmt.Println(u)
	}
}

type MockCrawler struct {
	mock.Mock
}

func (m *MockCrawler) GetUserList() ([]*User, error) {
	args := m.Called()
	return args.Get(0).([]*User), args.Error(1)
}

var MockUsers []*User

func init() {
	MockUsers = append(MockUsers, &User{"dj", 18})
	MockUsers = append(MockUsers, &User{"zhangsan", 20})
}

func TestGetUserList(t *testing.T) {
	crawler := new(MockCrawler)
	crawler.On("GetUserList").Return(MockUsers, nil)

	GetAndPrintUsers(crawler)

	crawler.AssertExpectations(t)
}
