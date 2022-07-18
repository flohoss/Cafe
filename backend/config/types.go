package config

type ErrorResponses uint

const (
	MissingInformation ErrorResponses = iota
	CannotCreate
	CannotUpdate
	CannotDelete
	CannotFind
	StillInUse
	CannotParse
)

func (e ErrorResponses) String() string {
	switch e {
	case MissingInformation:
		return "fehlende Informationen"
	case CannotCreate:
		return "kann nicht gespeichert werden"
	case CannotUpdate:
		return "kann nicht geändert werden"
	case CannotDelete:
		return "kann nicht gelöscht werden"
	case CannotFind:
		return "kann nicht gefunden werden"
	case StillInUse:
		return "noch in Verwendung"
	case CannotParse:
		return "kann nicht verarbeitet werden"
	default:
		return "unbekannt"
	}
}
