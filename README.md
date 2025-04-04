# haikube
Design, build, and deploy Haikube — a daily haiku generator using OpenAI, Go, and Kubernetes — with persistent storage and a simple user interface.

## Scaffolding

/haikube
│
├── /frontend
│   ├── /public
│   │   └── index.html           # Main HTML file
│   ├── /src
│   │   ├── /components
│   │   │   └── App.js           # React components (if you're using React)
│   │   ├── /assets
│   │   │   └── logo.png         # Images and other static assets
│   │   └── index.js             # Main JS file to load the React app
│   ├── package.json             # Frontend dependencies
│   └── webpack.config.js        # Webpack config for bundling (if using webpack)
│
├── /backend
│   ├── /cmd
│   │   └── haikube.go           # Main entry point for Go backend
│   ├── /pkg
│   │   ├── /handlers
│   │   │   └── haiku_handler.go  # Handler for generating and serving haikus
│   │   ├── /services
│   │   │   └── haiku_service.go  # Service that interacts with OpenAI API or generates haikus
│   │   ├── /models
│   │   │   └── haiku.go         # Data structure for a Haiku
│   │   └── /utils
│   │       └── logger.go        # Utility functions like logging
│   ├── go.mod                   # Go module dependencies
│   └── go.sum                   # Go module checksum
│
├── /config
│   └── config.yaml              # Configuration file for backend settings (e.g., OpenAI API key)
│
└── README.md                    # Project documentation
