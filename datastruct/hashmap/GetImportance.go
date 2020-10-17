/*
* @Author  :     knight
* @Date    :     2020/09/13 18:29:19
* @Email   :     knight2347@163.com
* @idea    :     leetcode 690 员工的重要性
 */

package hashmap

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

//DFS版
func getImportance(employees []*Employee, id int) int {
	importances := 0
	employeesMap := make(map[int]*Employee)
	for _, v := range employees {
		employeesMap[v.Id] = v
	}
	var dfs func(employee *Employee)
	dfs = func(employee *Employee) {
		importances += employee.Importance
		for _, v := range employee.Subordinates {
			dfs(employeesMap[v])
		}
	}
	dfs(employeesMap[id])
	return importances
}

//BFS版
func getImportance2(employees []*Employee, id int) int {
	importances := 0
	queue := make([]*Employee, 0)
	employeesMap := make(map[int]*Employee)
	for _, v := range employees {
		employeesMap[v.Id] = v
	}
	queue = append(queue, employeesMap[id])
	for len(queue) > 0 {
		top := queue[0]
		queue = queue[1:]
		importances += top.Importance
		for _, eId := range top.Subordinates {
			queue = append(queue, employeesMap[eId])
		}
	}
	return importances
}
