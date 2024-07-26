package client

type TokenReply struct {
	Token string `json:"token"`
}

type ListGamesRequest struct {
	Page         int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize     int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Search       string   `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	Statuses     []string `protobuf:"varint,4,rep,packed,name=statuses,proto3" json:"statuses,omitempty"`
	Categories   []string `protobuf:"varint,5,rep,packed,name=categories,proto3" json:"categories,omitempty"`
	Orientations []string `protobuf:"varint,6,rep,packed,name=orientations,proto3" json:"orientations,omitempty"`
	Labels       []string `protobuf:"varint,7,rep,packed,name=labels,proto3" json:"labels,omitempty"`
	Languages    []string `protobuf:"bytes,8,rep,name=languages,proto3" json:"languages,omitempty"`
}

type ListGamesReply struct {
	Records      []*PlayerGame `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	Page         int32         `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PageSize     int32         `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Total        int32         `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	HasPrevious  bool          `protobuf:"varint,5,opt,name=has_previous,json=hasPrevious,proto3" json:"has_previous,omitempty"`
	PreviousPage int32         `protobuf:"varint,6,opt,name=previous_page,json=previousPage,proto3" json:"previous_page,omitempty"`
	HasNext      bool          `protobuf:"varint,7,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
	NextPage     int32         `protobuf:"varint,8,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

type PlayerGame struct {
	Id           int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	GameId       string   `protobuf:"bytes,2,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	Url          string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Categories   []string `protobuf:"varint,4,rep,packed,name=categories,proto3" json:"categories,omitempty"`
	Orientations []string `protobuf:"varint,5,rep,packed,name=orientations,proto3" json:"orientations,omitempty"`
	Labels       []string `protobuf:"varint,6,rep,packed,name=labels,proto3" json:"labels,omitempty"`
	Assets       []*Asset `protobuf:"bytes,7,rep,name=assets,proto3" json:"assets,omitempty"`
	Status       string   `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	Version      string   `protobuf:"bytes,9,opt,name=version,proto3" json:"version,omitempty"`
	Weight       int32    `protobuf:"varint,54,opt,name=weight,proto3" json:"weight,omitempty"`
	CustomLabels []string `protobuf:"bytes,55,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty"`
}

type Asset struct {
	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Language    string `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Slug        string `protobuf:"bytes,4,opt,name=slug,proto3" json:"slug,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// For History
	Image_92X92 string `protobuf:"bytes,6,opt,name=image_92x92,json=image92x92,proto3" json:"image_92x92,omitempty"`
	// For Website Icon
	Image_264X264 string `protobuf:"bytes,7,opt,name=image_264x264,json=image264x264,proto3" json:"image_264x264,omitempty"`
	// For Mini Game Hub Icon
	Image_243X224 string `protobuf:"bytes,8,opt,name=image_243x224,json=image243x224,proto3" json:"image_243x224,omitempty"`
	// For Web Icon
	Image_320X468 string `protobuf:"bytes,9,opt,name=image_320x468,json=image320x468,proto3" json:"image_320x468,omitempty"`
	// For Launcher Horizontal Game Icon
	Image_395X481 string `protobuf:"bytes,10,opt,name=image_395x481,json=image395x481,proto3" json:"image_395x481,omitempty"`
	// For Launcher Vertical Game Icon
	Image_341X280 string `protobuf:"bytes,11,opt,name=image_341x280,json=image341x280,proto3" json:"image_341x280,omitempty"`
}
