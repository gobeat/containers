package containers

const (
	bindInvalidArgumentsError         = "binding error! Invalid arguments"
	resolveInvalidArgumentError       = "resolving error! Invalid arguments"
	injectInvalidTargetTypeError      = "injecting to %v is not supported"
	bindInvalidInstanceError          = "%v is not an instance of %v"
	bindInvalidConcreteError          = "non-supported kind of concrete. Got %v"
	bindInvalidStructConcreteError    = "expects %s. Got %s"
	resolveNotExistAbstractError      = "%v is not bound yet"
	resolveInvalidConcreteError       = "type %v is not supported"
	bindInvalidStructError            = "called structOf with a value that is not a pointer to a struct. (*MyStruct)(nil)"
	bindInvalidStructTypeError        = "called structOfType with a value that is not a pointer to a struct. (*MyStruct)(nil)"
	bindInvalidInterfaceError         = "called interfaceOf with a value that is not a pointer to an interface. (*MyInterface)(nil)"
	resolveInsufficientArgumentsError = "expects to have %v input arguments. Got %v"
	resolveNonValuesReturnedError     = "expects to have at least 1 value returned. Got 0"
)
