package modules

import (
    "fmt"
    "log"
    "strconv"

    "github.com/amarnathcjd/gogram/telegram"
    "AeonMusisBotGo/config"
    "AeonMusisBotGo/utils"
)

func InlineMusicSearch(i *telegram.InlineQuery) error {
    builder := i.Builder()
    q := i.Query

    log.Printf("Searching for query: '%s'", q)

    // Extract offset and chat index from the input offset
    var offset int32 = 0
    chatIndex := 0

    if i.Offset != "" {
        parts := strings.Split(i.Offset, ":")
        if len(parts) == 2 {
            chatIndex, _ = strconv.Atoi(parts[0])
            parsedOffset, _ := strconv.Atoi(parts[1])
            offset = int32(parsedOffset)
        }
    }

    // Search in all chats one by one
    var messages []telegram.NewMessage
    var err error

    for chatIndex < len(config.ChatList) {
        chatID := config.ChatList[chatIndex]
        messages, err = i.Client.GetMessages(
            chatID,
            &telegram.SearchOption{
                Limit:     20,
                AddOffset: offset,
                Filter:    &telegram.InputMessagesFilterMusic{},
                Query:     q,
            },
        )

        if err != nil {
            log.Printf("Error fetching messages from chat %d: %v", chatID, err)
            return err
        }

        if len(messages) > 0 {
            break // Stop searching if results are found
        }

        // Move to next chat and reset offset
        chatIndex++
        offset = 0
    }

    if len(messages) == 0 {
        // No results found in any chat
        builder.Article(
            "no_results",
            "No Results Found",
            "No matching audio messages found.",
            &telegram.ArticleOptions{
                Description: "Try a different search term",
                Thumb: telegram.InputWebDocument{
                    URL:      "https://github.com/5hojib/5hojib/raw/refs/heads/main/images/audio_thumb.png",
                    MimeType: "image/jpeg",
                },
            },
        )
    } else {
        for _, msg := range messages {
            audio := msg.Audio()
            ChatID := msg.Channel.ID
            MID := msg.ID
            if audio == nil {
                continue
            }

            var audioAttr *telegram.DocumentAttributeAudio
            for _, attr := range audio.Attributes {
                if a, ok := attr.(*telegram.DocumentAttributeAudio); ok {
                    audioAttr = a
                    break
                }
            }

            if audioAttr == nil {
                continue
            }

            title := audioAttr.Title
            if title == "" {
                title = "Untitled"
            }
            performer := audioAttr.Performer
            duration := audioAttr.Duration
            f_dur := utils.FormatDuration(duration)
            size := audio.Size
            f_size := utils.FormatSize(size)

            text := fmt.Sprintf(
                "Title: %s\nPerformer: %s\nDuration: %s",
                title,
                performer,
                f_dur,
            )

            var bt = telegram.Button
            builder.Article(
                title,
                fmt.Sprintf("%s\n%s | %s", performer, f_dur, f_size),
                text,
                &telegram.ArticleOptions{
                    ID: fmt.Sprintf("%d/%d", ChatID, MID),
                    ReplyMarkup: telegram.NewKeyboard().AddRow(
                        bt.SwitchInline("Search Again", true, ""),
                    ).Build(),
                    Thumb: telegram.InputWebDocument{
                        URL:      "https://github.com/5hojib/5hojib/raw/refs/heads/main/images/audio_thumb.png",
                        MimeType: "image/jpeg",
                    },
                },
            )
        }
    }

    // Prepare next offset: `chatIndex:offset`
    nextOffset := ""
    if len(messages) == 20 {
        nextOffset = fmt.Sprintf("%d:%d", chatIndex, offset+20)
    } else if chatIndex+1 < len(config.ChatList) {
        nextOffset = fmt.Sprintf("%d:0", chatIndex+1) // Move to the next chat
    }

    results := builder.Results()
    log.Printf("Sending %d results for query '%s'", len(results), q)

    _, err = i.Answer(results, telegram.InlineSendOptions{
        NextOffset: nextOffset,
        CacheTime:  0,
    })

    if err != nil {
        log.Printf("Error answering inline query: %v", err)
    }
    return err
}

func ChosenInlineHandler(u *telegram.InlineSend) error {
    a, b := utils.SplitAndFormat(u.ID)
    q := u.OriginalUpdate.Query
    msg, err := u.Client.GetMessages(
        a,
        &telegram.SearchOption{
            IDs: b,
        },
    )
    if err != nil {
        log.Printf("Error getting messages: %s", err)
    }
    if msg == nil {
        return nil
    }
    u.Edit("", &telegram.SendOptions{
        Media: msg[0].Media(),
        ReplyMarkup: telegram.NewKeyboard().AddRow(
            telegram.Button.SwitchInline(
                "Search Again",
                true, q,
            ),
        ).Build(),
    })
    return nil
}