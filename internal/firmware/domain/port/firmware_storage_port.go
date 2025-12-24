package port

type FirmwareStoragePort interface{
    Save(key string, file []byte) error
    Delete(key string) error
    GetFile(key string) ([]byte, error)
}