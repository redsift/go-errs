package errs

import "fmt"

//go:generate stringer -type=InternalState
//go:generate jsonenums -type=InternalState
//go:generate msgp -io=false

// Reference the values against https://github.com/redsift/guide/wiki/Error-Code-<InternalState>
type InternalState int

//msgp:ignore Verb Adjective Noun
type Verb int
type Adjective int
type Noun int

const lookupURLFormat = "https://github.com/redsift/guide/wiki/Error-Code-%s"

func (i InternalState) LookupURL() string {
	return fmt.Sprintf(lookupURLFormat, i)
}

// Message for the Code
func (i InternalState) Message() string {
	switch i {
	case Mochasippi:
		return "Service not available"
	case Breve:
		return "Webhook downstream error"
	case Papi:
		return "request object validation issue"
	case Instant:
		return "Can't parse JSON"
	case Turkish:
		return "Generic Aerospike Error"
	case Irish:
		return "Aerospike Connection Error"
	case Crema:
		return "Last Request Id out of sync"
	case Cappuccino:
		return "Parameter error"
	case Latte:
		return "Shutting down so action not permissible"
	case Flatwhite:
		return "Shutting down so action not permissible BUT when you sent the request the upstream server could not have known"
	case Melange:
		return "Not yet implemented"
	case Cortado:
		return "JWT token invalid"
	case Galao:
		return "JWT token valid but upstream invalid"
	case Kopisusu:
		return "HMAC Validation Failed"
	case Affogato:
		return "Static Configuration Error"
	case Macchiato:
		return "Try Again"
	case Bicerin:
		return "Service is experiencing High Traffic"
	case Bombón:
		return "Could not delete user data"
	case Mocha:
		return "Persist Failed Downstream"
	case Caphesuada:
		return "IMAP error"
	case Carajillo:
		return "JWT token valid but upstream invalid and not refreshable"
	case Espresso:
		return "Sift timeout"
	case Eiskaffee:
		return "JWT token expired"
	case Frappuccino:
		return "Job buried"
	case Iced:
		return "DAG node reports error"
	case Indianfilter:
		return "Archive Failed Upstream"
	case Kopiluwak:
		return "Client version has been blacklisted"
	case Kopitubruk:
		return "Nanomsg communication error"
	case Vienna:
		return "Could not create install data"
	case Yuanyang:
		return "Assert"
	case None:
		return "None"
	case Americano:
		return "JMAP error"
	case Cubano:
		return "User is not a Sift admin"
	case Zorro:
		return "External code error"
	case Doppio:
		return "Can't marshal JSON"
	case Romano:
		return "Inbox service reports error"
	case Guillermo:
		return "General Mongo Error"
	case Ristretto:
		return "No Github account for user in JWE"
	case Antoccino:
		return "API internal error" // API because message would be delivered to end-user on some branded domain.
	case Aulait:
		return "Not found"
	case Kula:
		return "Bender internal error"
	case Melya:
		return "Chronos internal error"
	case Marocchino:
		return "Dagger internal error"
	case Miel:
		return "Botfwk internal error"
	case Mazagran:
		return "slackstream internal error"
	case Palazzo:
		return "jmapstream internal error"
	case Medici:
		return "jmaparchive internal error"
	case Touba:
		return "DAG runtime error"
	default:
		return "Unknown"
	}
}

const (
	Mochasippi   InternalState = iota // Service not available
	Breve                             // Webhook upstream error
	Papi                              // request object validation issue
	Instant                           // Can't parse JSON
	Turkish                           // General Aerospike Exception
	Irish                             // Aerospike connection
	Crema                             // Last Request Id out of sync
	Cappuccino                        // Parameter error
	Unknown                           // Just WTF
	Latte                             // Shutting down so action not permissible
	Flatwhite                         // Shutting down so action not permissible BUT when you sent the request the upstream server could not have known. Basically the same as a latte but much cooler
	Melange                           // Not yet implemented
	Cortado                           // JWT token invalid
	Galao                             // JWT token valid but upstream invalid
	Kopisusu                          // HMAC validation error
	Affogato                          // Static configuration problem
	Macchiato                         // Try AGAIN
	Bicerin                           // Service is experiencing High Traffic
	Bombón                            // Could not delete user data
	Mocha                             // Persist Failed Downstream
	Caphesuada                        // IMAP error
	Carajillo                         // JWT token valid but upstream invalid and not refreshable
	Espresso                          // Sift timeout
	Eiskaffee                         // JWT token expired
	Frappuccino                       // Job buried
	Iced                              // DAG node reports error
	Indianfilter                      // Archive Failed Upstream
	Kopiluwak                         // Client version has been blacklisted
	Kopitubruk                        // Nanomsg communication error
	Vienna                            // Could not create install data
	Yuanyang                          // Assert
	None                              //
	Americano                         // JMAP error
	Cubano                            // User is not a Sift admin
	Zorro                             // External code error
	Doppio                            // Can't marshal JSON
	Romano                            // Inbox service reports error
	Guillermo                         // General Mongo Error
	Ristretto                         // No Github account for user in JWE
	Antoccino                         // Kerash (RPC/API) internal error
	Aulait                            // Not found
	Kula                              // Bender internal error
	Melya                             // Chronos internal error
	Marocchino                        // Dagger internal error
	Miel                              // Botfwk internal error
	Mazagran                          // slackstream internal error
	Palazzo                           // jmapstream internal error
	Medici                            // jmaparchive internal error
	Touba				  // DAG runtime error
	Pocillo
)
