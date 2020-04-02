package persistence_test

import (
	"image-compress-cli/pkg/domain/model"
	"image-compress-cli/pkg/infra/persistence"
	"testing"
)

// TestResizeJpgSuccess ResizeJpgの正常系テスト
func TestResizeJpgSuccess(t *testing.T) {
	ip := persistence.NewImagePersistence()
	originalFile := model.File{
		Dir:          "../../../tests/storages/before/",
		Extension:    "jpg",
		ExtLowerCase: "jpg",
		FileName:     "example1.jpg",
		Name:         "example1",
		Path:         "../../../tests/storages/before/example1.jpg",
	}

	t.Run("Jpgのリサイズをする", func(t *testing.T) {
		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example1.jpg", 960, 70,
		)
		if err != nil {
			t.Fatalf("\nfailed test %#v", err)
		}
	})

	t.Run("JPGのリサイズをする", func(t *testing.T) {
		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example2.jpg", 960, 70,
		)
		if err != nil {
			t.Fatalf("\nfailed test %#v", err)
		}
	})
}

// TestResizeJpgFailed ResizeJpgの異常系テスト
func TestResizeJpgFailed(t *testing.T) {
	ip := persistence.NewImagePersistence()

	t.Run("originalpathでファイルが存在しない", func(t *testing.T) {
		originalFile := model.File{
			Dir:          "../../../tests/storages/before",
			Extension:    "",
			ExtLowerCase: "",
			FileName:     "",
			Name:         "",
			Path:         "../../../tests/storages/before",
		}

		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example1.jpg", 960, 70,
		)
		if err == nil {
			t.Fatalf("\nfailed test")
		}
	})

	t.Run("originalpathでファイルが画像でない", func(t *testing.T) {
		originalFile := model.File{
			Dir:          "../../../tests/storages/before/",
			Extension:    "gitkeep",
			ExtLowerCase: "gitkeep",
			FileName:     ".gitkeep",
			Name:         "",
			Path:         "../../../tests/storages/before/.gitkeep",
		}

		_, err := ip.ResizeJpg(
			originalFile, "../../../tests/storages/after/", "example1.jpg", 960, 70,
		)
		if err == nil {
			t.Fatalf("\nfailed test")
		}
	})
}
