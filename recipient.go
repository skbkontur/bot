package bot

func (r recipient) Destination() string {
	return r.uid
}
