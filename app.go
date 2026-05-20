package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ollama/ollama/api"
)

// App struct
type App struct {
	ctx context.Context
}

/**
 * 翻訳したい言語をここに追加
 */
type Translator struct {
	Japanese string `json:"Japanese"`
}

/**
 * 使用するLLMを記載
 */
type config struct {
	Model string `json:"model_g3"`
}

// configファイルを読み込み
func (a *App) loadConfig() (*config, error) {
	f, err := os.Open("llm.json")
	if err != nil {
		log.Fatal("loadConfig os.Open err:", err)
		return nil, err
	}
	defer f.Close()

	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

/**
 * 20260516
 * HelloWorld.vueから文字列を受け取る
 */
func (a *App) ProcessMessage(msg string) string {
	fmt.Printf("受信メッセージ: %s\n", msg)

	var l_resTrans = a.Translate(msg)

	return "翻訳しました: " + l_resTrans
}

/**
 * ProcessMessageで受け取ったメッセージを
 * ここでollamaのAPIを呼び出して翻訳を行う
 */
func (a *App) Translate(a_msg string) string {
	cfg, err := a.loadConfig()
	if err != nil {
		log.Fatal("llm.jsonの読み込み失敗:", err)
		return ""
	}

	var modelName = cfg.Model

	client, err := api.ClientFromEnvironment()
	if err != nil {
		fmt.Printf("Ollama clientの生成に失敗: %v", err)
	}

	prompt := fmt.Sprintf(`
Translate the following English sentence to Japanese.

Output ONLY valid JSON in the format:
{"Japanese": "<translation here>"}

english: %s`, a_msg)

	req := &api.GenerateRequest{
		Model:  modelName,
		Prompt: prompt,
		//Mankind knew that they cannot change society. So instead of reflecting on themselves, they blamed the BeastsFormat: json.RawMessage(`{"key": "value"}`),
	}

	var l_sb strings.Builder

	err = client.Generate(context.Background(), req, func(resp api.GenerateResponse) error {
		l_sb.WriteString(resp.Response)
		return nil
	})

	if err != nil {
		fmt.Printf("生成失敗: %v", err)
		return ""
	}

	// 受信した JSON の出力
	var l_result Translator
	l_output := strings.TrimSpace(l_sb.String())

	//Markdownのコードブロックを除去する
	l_output = strings.TrimPrefix(l_output, "```json")
	l_output = strings.TrimPrefix(l_output, "```")
	l_output = strings.TrimSuffix(l_output, "```")
	l_output = strings.TrimSpace(l_output)

	if err := json.Unmarshal([]byte(l_output), &l_result); err != nil {
		fmt.Printf("JSONの出力失敗: %v\nRaw output:\n%s", err, l_output)
	}

	fmt.Printf("翻訳結果: %s\n", l_result)

	return l_result.Japanese
}
