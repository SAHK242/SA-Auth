package cache

//func NewUserCache(lifecycle fx.Lifecycle, props props, repository repository.UserRepository) c.MyCache[*ent.User] {
//	application := props.Config.MustGetString("APPLICATION")
//	ttl := props.Config.MustGetInt("REDIS_DEFAULT_TTL_MS")
//
//	keyEncoder := func(key string) string {
//		return fmt.Sprintf("%s:user:%s", application, key)
//	}
//
//	keyDecoder := func(key string) string {
//		return strings.Split(key, ":")[2]
//	}
//
//	errorHandler := func(err error) {
//		props.Logger.Errorln(err)
//	}
//
//	client := c.NewCache[*ent.User](props.Redis,
//		c.WithKeyGenerator[*ent.User](keyEncoder),
//		c.WithKeyDecoder[*ent.User](keyDecoder),
//		c.WithErrorHandler[*ent.User](errorHandler),
//		c.WithLoader[*ent.User](repository.FindById),
//		c.WithTTL[*ent.User](time.Duration(ttl)*time.Millisecond),
//	)
//
//	lifecycle.Append(fx.Hook{
//		OnStop: func(ctx context.Context) error {
//			client.Close()
//			return nil
//		},
//	})
//
//	return client
//}
