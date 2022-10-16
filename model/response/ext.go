package response

import "github.com/harrypunk/haikou_cloud/model"

func StudentToInfo(st *model.Student) *StudentInfo {
	return &StudentInfo{
		ID:     st.ID,
		Name:   st.Name,
		Age:    st.Age,
		Gender: st.Gender,
		Phone:  st.Phone,
	}
}

func StudentToDetail(st *model.Student) *StudentDetail {
	detail := StudentDetail{}

	detail.ID = st.ID
	detail.Name = st.Name
	detail.Age = st.Age
	detail.Gender = st.Gender
	detail.Phone = st.Phone

	return &detail
}
