package file

import (
	"fmt"
)

type FileRequest struct {
	Filename   string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Hash       string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Size       int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	UploadTime string `protobuf:"bytes,4,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"`
}

func (m *FileRequest) Reset() { *m = FileRequest{} }
func (m *FileRequest) String() string {
	return fmt.Sprintf("FileRequest{Filename:%q, Hash:%q, Size:%d, UploadTime:%q}", m.Filename, m.Hash, m.Size, m.UploadTime)
}
func (*FileRequest) ProtoMessage() {}
func (m *FileRequest) Marshal() ([]byte, error) {
	var buf []byte
	if m.Filename != "" {
		buf = append(buf, 0x0a)
		buf = append(buf, encodeVarint(uint64(len(m.Filename)))...)
		buf = append(buf, m.Filename...)
	}
	if m.Hash != "" {
		buf = append(buf, 0x12)
		buf = append(buf, encodeVarint(uint64(len(m.Hash)))...)
		buf = append(buf, m.Hash...)
	}
	if m.Size != 0 {
		buf = append(buf, 0x18)
		buf = append(buf, encodeVarint(uint64(m.Size))...)
	}
	if m.UploadTime != "" {
		buf = append(buf, 0x22)
		buf = append(buf, encodeVarint(uint64(len(m.UploadTime)))...)
		buf = append(buf, m.UploadTime...)
	}
	return buf, nil
}

func (m *FileRequest) Unmarshal(data []byte) error {
	*m = FileRequest{}
	i := 0
	for i < len(data) {
		fieldNumber, wireType, n := decodeTag(data[i:])
		i += n
		switch fieldNumber {
		case 1:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.Filename = string(data[i : i+int(strLen)])
			i += int(strLen)
		case 2:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.Hash = string(data[i : i+int(strLen)])
			i += int(strLen)
		case 3:
			val, n := decodeVarint(data[i:])
			i += n
			m.Size = int64(val)
		case 4:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.UploadTime = string(data[i : i+int(strLen)])
			i += int(strLen)
		default:
			switch wireType {
			case 0:
				_, n := decodeVarint(data[i:])
				i += n
			case 2:
				l, n := decodeVarint(data[i:])
				i += n + int(l)
			case 5:
				i += 8
			default:
				i++
			}
		}
	}
	return nil
}

type FileResponse struct {
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileId  int64  `protobuf:"varint,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
}

func (m *FileResponse) Reset() { *m = FileResponse{} }
func (m *FileResponse) String() string {
	return fmt.Sprintf("FileResponse{Success:%v, Message:%q, FileId:%d}", m.Success, m.Message, m.FileId)
}
func (*FileResponse) ProtoMessage() {}
func (m *FileResponse) Marshal() ([]byte, error) {
	var buf []byte
	if m.Success {
		buf = append(buf, 0x08, 0x01)
	}
	if m.Message != "" {
		buf = append(buf, 0x12)
		buf = append(buf, encodeVarint(uint64(len(m.Message)))...)
		buf = append(buf, m.Message...)
	}
	if m.FileId != 0 {
		buf = append(buf, 0x18)
		buf = append(buf, encodeVarint(uint64(m.FileId))...)
	}
	return buf, nil
}

func (m *FileResponse) Unmarshal(data []byte) error {
	*m = FileResponse{}
	i := 0
	for i < len(data) {
		fieldNumber, wireType, n := decodeTag(data[i:])
		i += n
		switch fieldNumber {
		case 1:
			m.Success = data[i] != 0
			i++
		case 2:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.Message = string(data[i : i+int(strLen)])
			i += int(strLen)
		case 3:
			val, n := decodeVarint(data[i:])
			i += n
			m.FileId = int64(val)
		default:
			switch wireType {
			case 0:
				_, n := decodeVarint(data[i:])
				i += n
			case 2:
				l, n := decodeVarint(data[i:])
				i += n + int(l)
			case 5:
				i += 8
			default:
				i++
			}
		}
	}
	return nil
}

type FileListRequest struct {
	Page     int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
}

func (m *FileListRequest) Reset() { *m = FileListRequest{} }
func (m *FileListRequest) String() string {
	return fmt.Sprintf("FileListRequest{Page:%d, PageSize:%d}", m.Page, m.PageSize)
}
func (*FileListRequest) ProtoMessage() {}
func (m *FileListRequest) Marshal() ([]byte, error) {
	var buf []byte
	if m.Page != 0 {
		buf = append(buf, 0x08)
		buf = append(buf, encodeVarint(uint64(m.Page))...)
	}
	if m.PageSize != 0 {
		buf = append(buf, 0x10)
		buf = append(buf, encodeVarint(uint64(m.PageSize))...)
	}
	return buf, nil
}

func (m *FileListRequest) Unmarshal(data []byte) error {
	*m = FileListRequest{}
	i := 0
	for i < len(data) {
		fieldNumber, wireType, n := decodeTag(data[i:])
		i += n
		switch fieldNumber {
		case 1:
			val, n := decodeVarint(data[i:])
			i += n
			m.Page = int32(val)
		case 2:
			val, n := decodeVarint(data[i:])
			i += n
			m.PageSize = int32(val)
		default:
			switch wireType {
			case 0:
				_, n := decodeVarint(data[i:])
				i += n
			case 2:
				l, n := decodeVarint(data[i:])
				i += n + int(l)
			case 5:
				i += 8
			default:
				i++
			}
		}
	}
	return nil
}

type FileRecord struct {
	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename   string `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Hash       string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Size       int64  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	UploadTime string `protobuf:"bytes,5,opt,name=upload_time,json=uploadTime,proto3" json:"upload_time,omitempty"`
}

func (m *FileRecord) Reset() { *m = FileRecord{} }
func (m *FileRecord) String() string {
	return fmt.Sprintf("FileRecord{Id:%d, Filename:%q, Hash:%q, Size:%d, UploadTime:%q}", m.Id, m.Filename, m.Hash, m.Size, m.UploadTime)
}
func (*FileRecord) ProtoMessage() {}
func (m *FileRecord) Marshal() ([]byte, error) {
	var buf []byte
	if m.Id != 0 {
		buf = append(buf, 0x08)
		buf = append(buf, encodeVarint(uint64(m.Id))...)
	}
	if m.Filename != "" {
		buf = append(buf, 0x12)
		buf = append(buf, encodeVarint(uint64(len(m.Filename)))...)
		buf = append(buf, m.Filename...)
	}
	if m.Hash != "" {
		buf = append(buf, 0x1a)
		buf = append(buf, encodeVarint(uint64(len(m.Hash)))...)
		buf = append(buf, m.Hash...)
	}
	if m.Size != 0 {
		buf = append(buf, 0x20)
		buf = append(buf, encodeVarint(uint64(m.Size))...)
	}
	if m.UploadTime != "" {
		buf = append(buf, 0x2a)
		buf = append(buf, encodeVarint(uint64(len(m.UploadTime)))...)
		buf = append(buf, m.UploadTime...)
	}
	return buf, nil
}

func (m *FileRecord) Unmarshal(data []byte) error {
	*m = FileRecord{}
	i := 0
	for i < len(data) {
		fieldNumber, wireType, n := decodeTag(data[i:])
		i += n
		switch fieldNumber {
		case 1:
			val, n := decodeVarint(data[i:])
			i += n
			m.Id = int64(val)
		case 2:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.Filename = string(data[i : i+int(strLen)])
			i += int(strLen)
		case 3:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.Hash = string(data[i : i+int(strLen)])
			i += int(strLen)
		case 4:
			val, n := decodeVarint(data[i:])
			i += n
			m.Size = int64(val)
		case 5:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.UploadTime = string(data[i : i+int(strLen)])
			i += int(strLen)
		default:
			switch wireType {
			case 0:
				_, n := decodeVarint(data[i:])
				i += n
			case 2:
				l, n := decodeVarint(data[i:])
				i += n + int(l)
			case 5:
				i += 8
			default:
				i++
			}
		}
	}
	return nil
}

type FileListResponse struct {
	Success bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string        `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Files   []*FileRecord `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
	Total   int64         `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (m *FileListResponse) Reset() { *m = FileListResponse{} }
func (m *FileListResponse) String() string {
	return fmt.Sprintf("FileListResponse{Success:%v, Message:%q, Files:%d, Total:%d}", m.Success, m.Message, len(m.Files), m.Total)
}
func (*FileListResponse) ProtoMessage() {}
func (m *FileListResponse) Marshal() ([]byte, error) {
	var buf []byte
	if m.Success {
		buf = append(buf, 0x08, 0x01)
	}
	if m.Message != "" {
		buf = append(buf, 0x12)
		buf = append(buf, encodeVarint(uint64(len(m.Message)))...)
		buf = append(buf, m.Message...)
	}
	for _, f := range m.Files {
		fData, _ := f.Marshal()
		buf = append(buf, 0x1a)
		buf = append(buf, encodeVarint(uint64(len(fData)))...)
		buf = append(buf, fData...)
	}
	if m.Total != 0 {
		buf = append(buf, 0x20)
		buf = append(buf, encodeVarint(uint64(m.Total))...)
	}
	return buf, nil
}

func (m *FileListResponse) Unmarshal(data []byte) error {
	*m = FileListResponse{}
	i := 0
	for i < len(data) {
		fieldNumber, wireType, n := decodeTag(data[i:])
		i += n
		switch fieldNumber {
		case 1:
			m.Success = data[i] != 0
			i++
		case 2:
			strLen, n := decodeVarint(data[i:])
			i += n
			m.Message = string(data[i : i+int(strLen)])
			i += int(strLen)
		case 3:
			lenVal, n := decodeVarint(data[i:])
			i += n
			f := new(FileRecord)
			f.Unmarshal(data[i : i+int(lenVal)])
			m.Files = append(m.Files, f)
			i += int(lenVal)
		case 4:
			val, n := decodeVarint(data[i:])
			i += n
			m.Total = int64(val)
		default:
			switch wireType {
			case 0:
				_, n := decodeVarint(data[i:])
				i += n
			case 2:
				l, n := decodeVarint(data[i:])
				i += n + int(l)
			case 5:
				i += 8
			default:
				i++
			}
		}
	}
	return nil
}

func encodeVarint(v uint64) []byte {
	var buf []byte
	for {
		b := byte(v & 0x7f)
		v >>= 7
		if v != 0 {
			buf = append(buf, b|0x80)
		} else {
			buf = append(buf, b)
			break
		}
	}
	return buf
}

func decodeVarint(buf []byte) (uint64, int) {
	var v uint64
	var shift uint
	var n int
	for _, b := range buf {
		n++
		v |= uint64(b&0x7f) << shift
		if b&0x80 == 0 {
			break
		}
		shift += 7
	}
	return v, n
}

func decodeTag(buf []byte) (fieldNumber, wireType uint64, n int) {
	v, n := decodeVarint(buf)
	fieldNumber = v >> 3
	wireType = v & 0x7
	return fieldNumber, wireType, n
}

func (m *FileRequest) XXX_DiscardUnknown()      {}
func (m *FileResponse) XXX_DiscardUnknown()     {}
func (m *FileListRequest) XXX_DiscardUnknown()  {}
func (m *FileRecord) XXX_DiscardUnknown()       {}
func (m *FileListResponse) XXX_DiscardUnknown() {}

func (m *FileRequest) XXX_Size() int {
	data, _ := m.Marshal()
	return len(data)
}

func (m *FileResponse) XXX_Size() int {
	data, _ := m.Marshal()
	return len(data)
}

func (m *FileListRequest) XXX_Size() int {
	data, _ := m.Marshal()
	return len(data)
}

func (m *FileRecord) XXX_Size() int {
	data, _ := m.Marshal()
	return len(data)
}

func (m *FileListResponse) XXX_Size() int {
	data, _ := m.Marshal()
	return len(data)
}
