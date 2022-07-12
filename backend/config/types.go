package config

type ErrorResponses uint

const (
	MissingInformation ErrorResponses = iota
	CannotCreate
	CannotDelete
	CannotFind
	StillInUse
)

func (e ErrorResponses) String() string {
	switch e {
	case MissingInformation:
		return "fehlende Informationen"
	case CannotCreate:
		return "kann nicht gespeichert werden"
	case CannotDelete:
		return "kann nicht gelöscht werden"
	case CannotFind:
		return "kann nicht gefunden werden"
	case StillInUse:
		return "noch in Verwendung"
	default:
		return "unbekannt"
	}
}
