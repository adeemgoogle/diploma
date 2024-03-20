package app

func (s *Server) Router() {
	s.App.Get("/getKey", s.GetKeys)
	s.App.Get("/h", s.HealthCheck)
}
