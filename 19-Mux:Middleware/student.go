package main

type Student struct {
	Id    string
	Name  string
	Grade int32
}

// return all data student
func GetStudents() []*Student {
	return students
}

// fungsi return data student with id selected
func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}

func init() {
	// make several data dummy to variable students
	students = append(students, &Student{Id: "s001", Name: "Akie", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "Kanao", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "Sagiri", Grade: 3})
}