// Copyright 2020 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package object

type Topic struct {
	Id            string `xorm:"varchar(100) notnull pk" json:"id"`
	Author        string `xorm:"varchar(100)" json:"author"`
	NodeId        string `xorm:"varchar(100)" json:"nodeId"`
	NodeName      string `xorm:"varchar(100)" json:"nodeName"`
	Title         string `xorm:"varchar(100)" json:"title"`
	CreatedTime   string `xorm:"varchar(100)" json:"createdTime"`
	LastReplyUser string `xorm:"varchar(100)" json:"lastReplyUser"`
	UpCount       int    `json:"upCount"`

	Content string `xorm:"mediumtext" json:"content"`
}

func GetTopics() []*Topic {
	topics := []*Topic{}
	err := adapter.engine.Asc("created_time").Find(&topics)
	if err != nil {
		panic(err)
	}

	return topics
}

func GetTopic(id string) *Topic {
	topic := Topic{Id: id}
	existed, err := adapter.engine.Get(&topic)
	if err != nil {
		panic(err)
	}

	if existed {
		return &topic
	} else {
		return nil
	}
}

func UpdateTopic(id string, topic *Topic) bool {
	if GetTopic(id) == nil {
		return false
	}

	_, err := adapter.engine.Id(id).AllCols().Update(topic)
	if err != nil {
		panic(err)
	}

	//return affected != 0
	return true
}

func AddTopic(topic *Topic) bool {
	affected, err := adapter.engine.Insert(topic)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteTopic(id string) bool {
	affected, err := adapter.engine.Id(id).Delete(&Topic{})
	if err != nil {
		panic(err)
	}

	return affected != 0
}