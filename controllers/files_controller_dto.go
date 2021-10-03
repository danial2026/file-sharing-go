package controllers

type FileType int

const (
	PROFILE_PIC FileType = iota
    POST_PIC
    Unknown
)

func FileTypFeromString(str string) FileType {
    switch str {
    case "PROFILE_PIC":
        return PROFILE_PIC
    case "POST_PIC":
        return POST_PIC
    }
    return Unknown
}

func (f FileType) Path() string {
	switch f {
    case PROFILE_PIC:
		return "downloadFile/users/"
    case POST_PIC:
		return "downloadFile/posts/"
    }
    return "Unknown"
}

type response struct {
    Timestamp string `json:"timestamp"`
    Status   int      `json:"status"`
    Message string `json:"message"`
    Path string `json:"path"`
}