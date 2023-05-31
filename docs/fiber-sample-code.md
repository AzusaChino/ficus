# Fiber Sample Codes

## Middleware

```go
func ApplyMiddleware(app *fiber.App) {

    // Match any route
    app.Use(func(c *fiber.Ctx) error {
        fmt.Println("🥇 First handler")
        return c.Next()
    })

    // Match all routes starting with /api
    app.Use("/api", func(c *fiber.Ctx) error {
        fmt.Println("🥈 Second handler")
        return c.Next()
    })

    // GET /api/list
    app.Get("/api/list", func(c *fiber.Ctx) error {
        fmt.Println("🥉 Last handler")
        return c.SendString("Hello, World 👋!")
    })

}
```

## Router

```go
func _(app *fiber.App) {

    // GET /api/register
    app.Get("/api/*", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("✋ %s", c.Params("*"))
        return c.SendString(msg) // => ✋ register
    })

    // GET /flights/LAX-SFO
    app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
        return c.SendString(msg) // => 💸 From: LAX, To: SFO
    })

    // GET /dictionary.txt
    app.Get("/:file.:ext", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
        return c.SendString(msg) // => 📃 dictionary.txt
    })

    // GET /john/75
    app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
        return c.SendString(msg) // => 👴 john is 75 years old
    })

    // GET /john
    app.Get("/:name", func(c *fiber.Ctx) error {
        msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
        return c.SendString(msg) // => Hello john 👋!
    })

    log.Fatal(app.Listen(":3000"))
}
```

## Serve Static

```go
func _(app *fiber.App) {
    app.Static("/", "./public")
    app.Static("/prefix", "./public")
    app.Static("/*", "./public/index.html")
}
```

## Websocket

```go
func _(app *fiber.App) {

    // 1. enable ws middleware
    app.Use("/ws", func(c *fiber.Ctx) error {
        // IsWebSocketUpgrade returns true if the client
        // requested upgrade to the WebSocket protocol.
        if websocket.IsWebSocketUpgrade(c) {
            c.Locals("allowed", true)
            return c.Next()
        }
        return fiber.ErrUpgradeRequired
    })

    // 2. sample ws usage
    app.Get("/ws/:id", websocket.New(func(conn *websocket.Conn) {
        log.Println(conn.Locals("allowed"))
        log.Println(conn.Params("id"))
        log.Println(conn.Query("v"))
        log.Println(conn.Cookies("session"))

        var (
            mt  int
            msg []byte
            err error
        )
        // for loop for listening
        for {
            if mt, msg, err = conn.ReadMessage(); err != nil {
                log.Fatalf("read error: %v\n", err)
            }
            log.Printf("recv msg: %s", msg)
            if err = conn.WriteMessage(mt, msg); err != nil {
                log.Fatalf("write error: %v\n", err)
            }
        }
    }))
}
```

## Render View

```go
type errorTemplateEngine struct{}

func (t errorTemplateEngine) Render(w io.Writer, name string, bind interface{}, layout ...string) error {
    return errors.New("errorTemplateEngine")
}

func (t errorTemplateEngine) Load() error { return nil }

func _(app *fiber.App) {

    app = fiber.New(fiber.Config{
        Views: errorTemplateEngine{},
    })

    app.Get("/", func(ctx *fiber.Ctx) error {
        return ctx.Render("home", fiber.Map{
            "title": "HomePage",
            "year":  2021,
        })
    })
}

```
