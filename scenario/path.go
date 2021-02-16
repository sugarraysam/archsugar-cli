package scenario

import "path"

func TasksDir(baseDir string) string {
	return path.Join(baseDir, "roles/main/tasks/master")
}

func VarsDir(baseDir string) string {
	return path.Join(baseDir, "roles/main/vars/master")
}

func EnabledDir(baseDir string) string {
	return path.Join(baseDir, "roles/main/tasks/master/enabled")
}
