/*
package internal
// 这个文件包含了录音相关的逻辑，使用 portaudio 库来实现从麦克风录音并保存为 WAV 文件的功能。
import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/gordonklaus/portaudio"
)

// RecordToFile 录制一段固定时间的音频用于 ASR 测试
func RecordToFile(seconds int, path string) error {
	portaudio.Initialize()
	defer portaudio.Terminate()

	const sampleRate = 16000
	const secondsToRecord = 2
	frames := make([]int16, sampleRate*seconds)

	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(frames), frames)
	if err != nil {
		return err
	}
	defer stream.Close()

	fmt.Println(">> 正在录音...")
	if err := stream.Start(); err != nil {
		return err
	}
	if err := stream.Read(); err != nil {
		return err
	}
	if err := stream.Stop(); err != nil {
		return err
	}

	return saveRawAsWav(path, frames, sampleRate)
}


func saveRawAsWav(path string, data []int16, sampleRate int) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = binary.Write(f, binary.LittleEndian, data)
	return err
}
*/

package internal
