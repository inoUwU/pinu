package port

// SSEBroker SSEメッセージブローカーのポート
type SSEBroker interface {
	// Publish メッセージをキューに追加する
	Publish(msg string)

	// Consume キューからメッセージを取り出す
	// 存在しない場合は false を返す
	Consume() (string, bool)
}
