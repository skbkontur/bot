package bot

import "fmt"

// Dialog returns Dialog struct by ID
func (r *Room) Dialog(id int64) *Dialog {
	log.WithFields(r).Debug(fmt.Sprintf("Dialog [%d] queried", id))
	d, ok := r.dialogs[id]
	if !ok {
		d = &Dialog{
			ID:        id,
			Room:      r,
			moderator: r.moderator.ModeratorID(),
		}
		r.dialogs[id] = d
	}
	return d
}
