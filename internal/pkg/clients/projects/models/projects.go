package models

type ProjectDto struct {
	Name          string `json:"name" binding:"required"`
	RepositoryURL string `json:"repositoryURL" binding:"required"`
}	

type GetProjectDto struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	LastCheckDate string `json:"lastCheckDate"`
	RepositoryURL string `json:"repositoryURL"`
} 

type ProjectGroupDto struct {
	Name  string   `json:"name" binding:"required"`
	Items []string `json:"items"`
} 
