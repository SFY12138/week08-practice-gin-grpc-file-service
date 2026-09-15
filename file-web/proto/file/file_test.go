package file

import (
	"testing"
)

func TestFileRequestMarshalUnmarshal(t *testing.T) {
	req := &FileRequest{
		Filename:   "test.txt",
		Hash:       "abc123",
		Size:       1024,
		UploadTime: "2024-01-01T00:00:00Z",
	}

	data, err := req.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	req2 := &FileRequest{}
	err = req2.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if req2.Filename != req.Filename {
		t.Errorf("Filename mismatch: got %q, want %q", req2.Filename, req.Filename)
	}
	if req2.Hash != req.Hash {
		t.Errorf("Hash mismatch: got %q, want %q", req2.Hash, req.Hash)
	}
	if req2.Size != req.Size {
		t.Errorf("Size mismatch: got %d, want %d", req2.Size, req.Size)
	}
	if req2.UploadTime != req.UploadTime {
		t.Errorf("UploadTime mismatch: got %q, want %q", req2.UploadTime, req.UploadTime)
	}
}

func TestFileResponseMarshalUnmarshal(t *testing.T) {
	resp := &FileResponse{
		Success: true,
		Message: "Success",
		FileId:  123,
	}

	data, err := resp.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	resp2 := &FileResponse{}
	err = resp2.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if resp2.Success != resp.Success {
		t.Errorf("Success mismatch: got %v, want %v", resp2.Success, resp.Success)
	}
	if resp2.Message != resp.Message {
		t.Errorf("Message mismatch: got %q, want %q", resp2.Message, resp.Message)
	}
	if resp2.FileId != resp.FileId {
		t.Errorf("FileId mismatch: got %d, want %d", resp2.FileId, resp.FileId)
	}
}

func TestFileListRequestMarshalUnmarshal(t *testing.T) {
	req := &FileListRequest{
		Page:     2,
		PageSize: 20,
	}

	data, err := req.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	req2 := &FileListRequest{}
	err = req2.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if req2.Page != req.Page {
		t.Errorf("Page mismatch: got %d, want %d", req2.Page, req.Page)
	}
	if req2.PageSize != req.PageSize {
		t.Errorf("PageSize mismatch: got %d, want %d", req2.PageSize, req.PageSize)
	}
}

func TestFileRecordMarshalUnmarshal(t *testing.T) {
	record := &FileRecord{
		Id:         1,
		Filename:   "test.txt",
		Hash:       "abc123",
		Size:       1024,
		UploadTime: "2024-01-01T00:00:00Z",
	}

	data, err := record.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	record2 := &FileRecord{}
	err = record2.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if record2.Id != record.Id {
		t.Errorf("Id mismatch: got %d, want %d", record2.Id, record.Id)
	}
	if record2.Filename != record.Filename {
		t.Errorf("Filename mismatch: got %q, want %q", record2.Filename, record.Filename)
	}
}

func TestFileListResponseMarshalUnmarshal(t *testing.T) {
	resp := &FileListResponse{
		Success: true,
		Message: "Success",
		Files: []*FileRecord{
			{
				Id:         1,
				Filename:   "test1.txt",
				Hash:       "hash1",
				Size:       100,
				UploadTime: "2024-01-01T00:00:00Z",
			},
			{
				Id:         2,
				Filename:   "test2.txt",
				Hash:       "hash2",
				Size:       200,
				UploadTime: "2024-01-02T00:00:00Z",
			},
		},
		Total: 2,
	}

	data, err := resp.Marshal()
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	resp2 := &FileListResponse{}
	err = resp2.Unmarshal(data)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if resp2.Success != resp.Success {
		t.Errorf("Success mismatch: got %v, want %v", resp2.Success, resp.Success)
	}
	if len(resp2.Files) != len(resp.Files) {
		t.Errorf("Files length mismatch: got %d, want %d", len(resp2.Files), len(resp.Files))
	}
	if resp2.Total != resp.Total {
		t.Errorf("Total mismatch: got %d, want %d", resp2.Total, resp.Total)
	}
}

func TestVarintEncoding(t *testing.T) {
	testCases := []struct {
		input    uint64
		expected []byte
	}{
		{0, []byte{0x00}},
		{1, []byte{0x01}},
		{127, []byte{0x7f}},
		{128, []byte{0x80, 0x01}},
		{256, []byte{0x80, 0x02}},
		{300, []byte{0xac, 0x02}},
	}

	for _, tc := range testCases {
		encoded := encodeVarint(tc.input)
		if len(encoded) != len(tc.expected) {
			t.Errorf("encodeVarint(%d): length mismatch: got %d bytes, want %d bytes", tc.input, len(encoded), len(tc.expected))
			continue
		}
		for i := range encoded {
			if encoded[i] != tc.expected[i] {
				t.Errorf("encodeVarint(%d): byte %d mismatch: got 0x%x, want 0x%x", tc.input, i, encoded[i], tc.expected[i])
			}
		}

		decoded, n := decodeVarint(tc.expected)
		if decoded != tc.input {
			t.Errorf("decodeVarint(%v): got %d, want %d", tc.expected, decoded, tc.input)
		}
		if n != len(tc.expected) {
			t.Errorf("decodeVarint(%v): consumed %d bytes, expected %d", tc.expected, n, len(tc.expected))
		}
	}
}
