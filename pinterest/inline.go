package pinterest

import (
        "fmt"
        "strings"
        "math/rand"

        "github.com/PaulSonOfLars/gotgbot/v2"
        "github.com/PaulSonOfLars/gotgbot/v2/ext"
        "github.com/Mishel-07/PinterestBot/settings"
)

func FindImageInline(b *gotgbot.Bot, ctx *ext.Context) error {                
        var query string
        var caption string 
        if strings.Contains(ctx.InlineQuery.Query, "!cap") {
                split := strings.Split(ctx.InlineQuery.Query, "!cap")    
                caption = split[1]
                if split[0] != "" {
                        query = split[0]
                } else {
                        query = ctx.InlineQuery.Query
                }
        } else {
                query = ctx.InlineQuery.Query
        }
        if query == "" {
                _, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
                        gotgbot.InlineQueryResultArticle{
                                Id: fmt.Sprintf("%d", rand.Int()),
                                Title: "No Query Provided",
                                InputMessageContent: &gotgbot.InputTextMessageContent{
                                        MessageText: "Please provide a query.",
                                },
                        },
                }, nil)
                return err
        }
        quotequery := strings.Replace(query, " ", "+", -1)
        urls, err := settings.SearchBingInline(quotequery)
        if err != nil {
                fmt.Println(err)
                _, err = ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
                        gotgbot.InlineQueryResultArticle{
                                Id: fmt.Sprintf("%d", rand.Int()),
                                Title: "Image not found",
                                InputMessageContent: &gotgbot.InputTextMessageContent{
                                        MessageText: "Image not found for your query.",
                                },
                        },
            }, nil)
            return err
        }

        media := make([]gotgbot.InlineQueryResult, 0)
        for _, item := range urls.Result {
                if item.URL != "" {
                        if caption != "" {
                                media = append(media, gotgbot.InlineQueryResultPhoto{
                                        Id: fmt.Sprintf("%d", rand.Int()),
                                        PhotoUrl: item.URL,   
                                        Caption: caption,
                                        Title: "Found Image",
                                        ThumbnailUrl: item.URL,
                                })
                        } else {
                                media = append(media, gotgbot.InlineQueryResultPhoto{
                                        Id: fmt.Sprintf("%d", rand.Int()),
                                        PhotoUrl: item.URL,                                                                             
                                        Title: "Found Image",
                                        ThumbnailUrl: item.URL,
                                })
                        }
                } else {
                        fmt.Println("Skipped empty URL")
                }
        }

        if len(media) == 0 {
                _, err := ctx.InlineQuery.Answer(b, []gotgbot.InlineQueryResult{
                        gotgbot.InlineQueryResultArticle{
                                Title: "No Images Found",
                                InputMessageContent: &gotgbot.InputTextMessageContent{
                                        MessageText: "No images found for your query.",
                                },
                        },
                }, nil)
                return err
        }

            
        _, err = ctx.InlineQuery.Answer(b, media, &gotgbot.AnswerInlineQueryOpts{})
        return err
}
