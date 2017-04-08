package environment

func (env *Env) Check(err error) bool {
	if err != nil {
		env.Log.Debug(err)
		return true
	}
	return false
}
