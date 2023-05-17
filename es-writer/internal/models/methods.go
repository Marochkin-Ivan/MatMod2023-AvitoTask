package models

func (f FileInfo) ToESDoc() ESDoc {
	return ESDoc{
		Title:        f.Name,
		Requirements: f.PositionRequirements,
		Keywords:     f.ProfessionalSphereName,
		Salary:       float64(f.SalaryMin),
		Region:       f.RegionName,
		CompanyName:  f.FullCompanyName,
		Schedule:     f.ScheduleType,
		Experience:   f.RequiredExperience,
		Employment:   f.BusyType,
		CreatedAt:    f.DateCreate,
	}
}
