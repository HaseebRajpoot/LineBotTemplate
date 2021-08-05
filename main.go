// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
	# Line Bot Environment initialization
MAIN_UID_OLD = 'u6b4ec5b657f07cf28612b2be96661f68'
MAIN_UID = 'u6b4ec5b657f07cf28612b2be96661f68'
main_silent = False
administrator = os.getenv('ADMIN', None)
group_admin = os.getenv('G_ADMIN', None)
group_mod = os.getenv('G_MOD', None)
if administrator is None:
    print('The SHA224 of ADMIN not defined. Program will be terminated.')
    sys.exit(1)
if group_admin is None:
    print('The SHA224 of G_ADMIN not defined. Program will be terminated.')
    sys.exit(1)
if group_mod is None:
    print('The SHA224 of G_MOD not defined. Program will be terminated.')
    sys.exit(1)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

@handler.add(JoinEvent)
def handle_join(event):
    src = event.source
    reply_token = event.reply_token
    cid = line_api_proc.source_channel_id(src)

    global command_executor

    if not isinstance(event.source, SourceUser):
        group_data = gb.get_group_by_id(cid)
        if group_data is None:
            added = gb.new_data(cid, MAIN_UID, administrator)
            msg_track.new_data(cid)

            api_reply(reply_token, 
                      [TextMessage(text=u'群組資料註冊{}。'.format(u'成功' if added else u'失敗')),
                       introduction_template()], 
                       cid)
        else:
            api_reply(reply_token, 
                      [TextMessage(text=u'群組資料已存在。'),
                       TextMessage(text=command_exec.G(src, [None, None, None])),
                       introduction_template()], 
                       cid)
	
		}
	}
}
