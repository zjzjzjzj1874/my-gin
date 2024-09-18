package model

type Account struct {
	Name string `json:"name"` // 账户名称
	Age  int    `json:"age"`  // 年龄
	Id   int    `json:"id"`   // id
}

func List() ([]Account, error) {
	return []Account{
		{
			Name: "张三",
			Age:  18,
			Id:   1,
		},
		{
			Name: "李四",
			Age:  20,
			Id:   2,
		},
	}, nil
}
