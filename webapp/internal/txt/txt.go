// package txt - texts constants and i18n
package txt

const (
	ErrCantDetectStorageTypeForDSN = "не удалось определить тип хранилища для '%s'"
)

// T - выполняет перевод текста согласно локали
func T(_s string) string {
	return _s
}
