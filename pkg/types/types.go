/*
@Description: just go
@Author: skipper
@Date: 2020/1/13
@Time: 3:50 PM
@ProjectName fileUpdater
*/
package types


type Updater struct {
	FilePath string
	PostUpdateHook []Hook
	PreUpdateHook []Hook
}

func (u Updater)GetFileContent()(content []byte,err error)  {

	return nil, nil
}
func (u Updater)UpdateFile(date []byte) error  {
	return nil
}

func (u Updater)execPreHook()  {

}
func (u Updater)execPostHook()  {

}


type Hook interface {
	Do() error
}

type CommandHook struct {
	Commands []string
	Mode     string
}

func (c CommandHook)Do() error {
	return nil
}