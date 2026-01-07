package main

import "sync"

var (
	languageSync           sync.Map
	fileNameLanguageSync   sync.Map
	hiddenFileLanguageSync sync.Map
)

func init() {
	for k, v := range Languages {
		languageSync.Store(k, v)
	}

	for k, v := range FilenameLanguages {
		fileNameLanguageSync.Store(k, v)
	}

	for k, v := range HiddenFileLanguages {
		hiddenFileLanguageSync.Store(k, v)
	}

}

// Language represents a programming language with its comment patterns
type Language struct {
	Name              string
	Extensions        []string
	SingleLineComment string
	MultiLineStart    string
	MultiLineEnd      string
	StringDelimiters  []string
	NestedComments    bool
}

// Languages defines all supported programming languages and their comment patterns
var Languages = func() map[string]*Language {
	const capacity = 115
	m := make(map[string]*Language, capacity)

	m[".go"] = &Language{
		Name:              "Go",
		Extensions:        []string{".go"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "`"},
	}
	m[".js"] = &Language{
		Name:              "JavaScript",
		Extensions:        []string{".js"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'", "`"},
	}
	m[".ts"] = &Language{
		Name:              "TypeScript",
		Extensions:        []string{".ts"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'", "`"},
	}
	m[".tsx"] = &Language{
		Name:              "TypeScript JSX",
		Extensions:        []string{".tsx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".jsx"] = &Language{
		Name:              "JavaScript JSX",
		Extensions:        []string{".jsx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".html"] = &Language{
		Name:              "HTML",
		Extensions:        []string{".html", ".htm"},
		SingleLineComment: "",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	}
	m[".htm"] = &Language{
		Name:              "HTML",
		Extensions:        []string{".html", ".htm"},
		SingleLineComment: "",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	}
	m[".py"] = &Language{
		Name:              "Python",
		Extensions:        []string{".py"},
		SingleLineComment: "#",
		MultiLineStart:    `"""`,
		MultiLineEnd:      `"""`,
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".rb"] = &Language{
		Name:              "Ruby",
		Extensions:        []string{".rb"},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".java"] = &Language{
		Name:              "Java",
		Extensions:        []string{".java"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".c"] = &Language{
		Name:              "C",
		Extensions:        []string{".c"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".h"] = &Language{
		Name:              "C Header",
		Extensions:        []string{".h"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".cpp"] = &Language{
		Name:              "C++",
		Extensions:        []string{".cpp", ".cc", ".cxx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".cc"] = &Language{
		Name:              "C++",
		Extensions:        []string{".cpp", ".cc", ".cxx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".hpp"] = &Language{
		Name:              "C++ Header",
		Extensions:        []string{".hpp"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".cs"] = &Language{
		Name:              "C#",
		Extensions:        []string{".cs"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".php"] = &Language{
		Name:              "PHP",
		Extensions:        []string{".php"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".swift"] = &Language{
		Name:              "Swift",
		Extensions:        []string{".swift"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\""},
		NestedComments:    true,
	}
	m[".kt"] = &Language{
		Name:              "Kotlin",
		Extensions:        []string{".kt"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\""},
		NestedComments:    true,
	}
	m[".rs"] = &Language{
		Name:              "Rust",
		Extensions:        []string{".rs"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\""},
		NestedComments:    true,
	}
	m[".scala"] = &Language{
		Name:              "Scala",
		Extensions:        []string{".scala"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".json"] = &Language{
		Name:              "JSON",
		Extensions:        []string{".json"},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".yaml"] = &Language{
		Name:              "YAML",
		Extensions:        []string{".yaml", ".yml"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".yml"] = &Language{
		Name:              "YAML",
		Extensions:        []string{".yaml", ".yml"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".md"] = &Language{
		Name:              "Markdown",
		Extensions:        []string{".md"},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".css"] = &Language{
		Name:              "CSS",
		Extensions:        []string{".css"},
		SingleLineComment: "",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".scss"] = &Language{
		Name:              "SCSS",
		Extensions:        []string{".scss"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".sass"] = &Language{
		Name:              "Sass",
		Extensions:        []string{".sass"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".less"] = &Language{
		Name:              "Less",
		Extensions:        []string{".less"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".sql"] = &Language{
		Name:              "SQL",
		Extensions:        []string{".sql"},
		SingleLineComment: "--",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".sh"] = &Language{
		Name:              "Shell",
		Extensions:        []string{".sh", ".bash"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".bash"] = &Language{
		Name:              "Shell",
		Extensions:        []string{".sh", ".bash"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".xml"] = &Language{
		Name:              "XML",
		Extensions:        []string{".xml"},
		SingleLineComment: "",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	}
	m[".vue"] = &Language{
		Name:              "Vue",
		Extensions:        []string{".vue"},
		SingleLineComment: "//",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	}
	m[".svelte"] = &Language{
		Name:              "Svelte",
		Extensions:        []string{".svelte"},
		SingleLineComment: "//",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	}
	m[".lua"] = &Language{
		Name:              "Lua",
		Extensions:        []string{".lua"},
		SingleLineComment: "--",
		MultiLineStart:    "--[[",
		MultiLineEnd:      "]]",
	}
	m[".r"] = &Language{
		Name:              "R",
		Extensions:        []string{".r", ".R"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".R"] = &Language{
		Name:              "R",
		Extensions:        []string{".r", ".R"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".pl"] = &Language{
		Name:              "Perl",
		Extensions:        []string{".pl", ".pm"},
		SingleLineComment: "#",
		MultiLineStart:    "=pod",
		MultiLineEnd:      "=cut",
	}
	m[".pm"] = &Language{
		Name:              "Perl",
		Extensions:        []string{".pl", ".pm"},
		SingleLineComment: "#",
		MultiLineStart:    "=pod",
		MultiLineEnd:      "=cut",
	}
	m[".ex"] = &Language{
		Name:              "Elixir",
		Extensions:        []string{".ex", ".exs"},
		SingleLineComment: "#",
		MultiLineStart:    `"""`,
		MultiLineEnd:      `"""`,
	}
	m[".exs"] = &Language{
		Name:              "Elixir",
		Extensions:        []string{".ex", ".exs"},
		SingleLineComment: "#",
		MultiLineStart:    `"""`,
		MultiLineEnd:      `"""`,
	}
	m[".erl"] = &Language{
		Name:              "Erlang",
		Extensions:        []string{".erl"},
		SingleLineComment: "%",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".hs"] = &Language{
		Name:              "Haskell",
		Extensions:        []string{".hs"},
		SingleLineComment: "--",
		MultiLineStart:    "{-",
		MultiLineEnd:      "-}",
	}
	m[".clj"] = &Language{
		Name:              "Clojure",
		Extensions:        []string{".clj"},
		SingleLineComment: ";",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".toml"] = &Language{
		Name:              "TOML",
		Extensions:        []string{".toml"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".ini"] = &Language{
		Name:              "INI",
		Extensions:        []string{".ini"},
		SingleLineComment: ";",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".dockerfile"] = &Language{
		Name:              "Dockerfile",
		Extensions:        []string{".dockerfile"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".makefile"] = &Language{
		Name:              "Makefile",
		Extensions:        []string{".makefile"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".tf"] = &Language{
		Name:              "Terraform",
		Extensions:        []string{".tf"},
		SingleLineComment: "#",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".proto"] = &Language{
		Name:              "Protocol Buffers",
		Extensions:        []string{".proto"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".graphql"] = &Language{
		Name:              "GraphQL",
		Extensions:        []string{".graphql", ".gql"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".gql"] = &Language{
		Name:              "GraphQL",
		Extensions:        []string{".graphql", ".gql"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".txt"] = &Language{
		Name:              "Text",
		Extensions:        []string{".txt"},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".hcl"] = &Language{
		Name:              "HCL",
		Extensions:        []string{".hcl"},
		SingleLineComment: "#",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".y"] = &Language{
		Name:              "Yacc",
		Extensions:        []string{".y"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".nix"] = &Language{
		Name:              "Nix",
		Extensions:        []string{".nix"},
		SingleLineComment: "#",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".json5"] = &Language{
		Name:              "JSON5",
		Extensions:        []string{".json5"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".s"] = &Language{
		Name:              "Assembly",
		Extensions:        []string{".s", ".S", ".asm"},
		SingleLineComment: ";",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".S"] = &Language{
		Name:              "Assembly",
		Extensions:        []string{".s", ".S", ".asm"},
		SingleLineComment: ";",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".asm"] = &Language{
		Name:              "Assembly",
		Extensions:        []string{".s", ".S", ".asm"},
		SingleLineComment: ";",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".dart"] = &Language{
		Name:              "Dart",
		Extensions:        []string{".dart"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".groovy"] = &Language{
		Name:              "Groovy",
		Extensions:        []string{".groovy"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".jl"] = &Language{
		Name:              "Julia",
		Extensions:        []string{".jl"},
		SingleLineComment: "#",
		MultiLineStart:    "#=",
		MultiLineEnd:      "=#",
		StringDelimiters:  []string{"\"", "'"},
	}
	m[".coffee"] = &Language{
		Name:              "CoffeeScript",
		Extensions:        []string{".coffee"},
		SingleLineComment: "#",
		MultiLineStart:    "###",
		MultiLineEnd:      "###",
		StringDelimiters:  []string{"\"", "'", "`"},
	}
	m[".pug"] = &Language{
		Name:              "Pug",
		Extensions:        []string{".pug"},
		SingleLineComment: "//",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".jade"] = &Language{
		Name:              "Jade",
		Extensions:        []string{".jade"},
		SingleLineComment: "//",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".twig"] = &Language{
		Name:              "Twig",
		Extensions:        []string{".twig"},
		SingleLineComment: "{#",
		MultiLineStart:    "{#",
		MultiLineEnd:      "#}",
	}
	m[".ejs"] = &Language{
		Name:              "EJS",
		Extensions:        []string{".ejs"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".haml"] = &Language{
		Name:              "Haml",
		Extensions:        []string{".haml"},
		SingleLineComment: "-#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".tsv"] = &Language{
		Name:              "TSV",
		Extensions:        []string{".tsv"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".csv"] = &Language{
		Name:              "CSV",
		Extensions:        []string{".csv"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".awk"] = &Language{
		Name:              "AWK",
		Extensions:        []string{".awk"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".sed"] = &Language{
		Name:              "Sed",
		Extensions:        []string{".sed"},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".v"] = &Language{
		Name:              "V",
		Extensions:        []string{".v"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".zig"] = &Language{
		Name:              "Zig",
		Extensions:        []string{".zig"},
		SingleLineComment: "//",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".ada"] = &Language{
		Name:              "Ada",
		Extensions:        []string{".ada", ".adb", ".ads"},
		SingleLineComment: "--",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".adb"] = &Language{
		Name:              "Ada",
		Extensions:        []string{".ada", ".adb", ".ads"},
		SingleLineComment: "--",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".ml"] = &Language{
		Name:              "OCaml",
		Extensions:        []string{".ml"},
		SingleLineComment: "(*",
		MultiLineEnd:      "*)",
		MultiLineStart:    "(*",
	}
	m[".mli"] = &Language{
		Name:              "OCaml Interface",
		Extensions:        []string{".mli"},
		SingleLineComment: "(*",
		MultiLineEnd:      "*)",
		MultiLineStart:    "(*",
	}
	m[".fs"] = &Language{
		Name:              "F#",
		Extensions:        []string{".fs", ".fsx", ".fsi"},
		SingleLineComment: "//",
		MultiLineStart:    "(*",
		MultiLineEnd:      "*)",
	}
	m[".fsx"] = &Language{
		Name:              "F# Script",
		Extensions:        []string{".fs", ".fsx", ".fsi"},
		SingleLineComment: "//",
		MultiLineStart:    "(*",
		MultiLineEnd:      "*)",
	}
	m[".nim"] = &Language{
		Name:              "Nim",
		Extensions:        []string{".nim"},
		SingleLineComment: "#",
		MultiLineStart:    "#[",
		MultiLineEnd:      "]#",
	}
	m[".d"] = &Language{
		Name:              "D",
		Extensions:        []string{".d"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".pas"] = &Language{
		Name:              "Pascal",
		Extensions:        []string{".pas"},
		SingleLineComment: "//",
		MultiLineStart:    "(*",
		MultiLineEnd:      "*)",
	}
	m[".vb"] = &Language{
		Name:              "Visual Basic",
		Extensions:        []string{".vb"},
		SingleLineComment: "'",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".vbs"] = &Language{
		Name:              "VBScript",
		Extensions:        []string{".vbs"},
		SingleLineComment: "'",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m[".ps1"] = &Language{
		Name:              "PowerShell",
		Extensions:        []string{".ps1"},
		SingleLineComment: "#",
		MultiLineStart:    "<#",
		MultiLineEnd:      "#>",
	}
	m[".asmx"] = &Language{
		Name:              "ASP.NET",
		Extensions:        []string{".asmx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".aspx"] = &Language{
		Name:              "ASP.NET",
		Extensions:        []string{".aspx"},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m[".cshtml"] = &Language{
		Name:              "Razor",
		Extensions:        []string{".cshtml"},
		SingleLineComment: "@*",
		MultiLineStart:    "@*",
		MultiLineEnd:      "*@",
	}
	m[".vbhtml"] = &Language{
		Name:              "Razor VB",
		Extensions:        []string{".vbhtml"},
		SingleLineComment: "@*",
		MultiLineStart:    "@*",
		MultiLineEnd:      "*@",
	}

	return m
}()

// BinaryExtensions contains file extensions that should be skipped
var BinaryExtensions = map[string]bool{
	// Images
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".bmp":  true,
	".ico":  true,
	".svg":  true,
	".webp": true,
	".tiff": true,
	".tif":  true,
	".psd":  true,
	".raw":  true,
	".heif": true,
	".heic": true,
	".avif": true,

	// Documents
	".pdf":  true,
	".doc":  true,
	".docx": true,
	".xls":  true,
	".xlsx": true,
	".ppt":  true,
	".pptx": true,
	".odt":  true,
	".ods":  true,
	".odp":  true,
	".rtf":  true,
	".epub": true,
	".mobi": true,

	// Archives
	".zip":  true,
	".tar":  true,
	".gz":   true,
	".tgz":  true,
	".bz2":  true,
	".xz":   true,
	".lz":   true,
	".lzma": true,
	".rar":  true,
	".7z":   true,
	".cab":  true,
	".iso":  true,
	".dmg":  true,
	".deb":  true,
	".rpm":  true,
	".apk":  true,
	".msi":  true,

	// Executables and libraries
	".exe":   true,
	".dll":   true,
	".so":    true,
	".dylib": true,
	".bin":   true,
	".com":   true,
	".mach":  true,
	".elf":   true,

	// Data files
	".dat":     true,
	".db":      true,
	".sqlite":  true,
	".sqlite3": true,
	".mdb":     true,
	".accdb":   true,
	".frm":     true,
	".ibd":     true,
	".dbf":     true,

	// Audio
	".mp3":  true,
	".wav":  true,
	".flac": true,
	".ogg":  true,
	".aac":  true,
	".wma":  true,
	".m4a":  true,
	".aiff": true,
	".mid":  true,
	".midi": true,

	// Video
	".mp4":  true,
	".avi":  true,
	".mov":  true,
	".wmv":  true,
	".flv":  true,
	".mkv":  true,
	".webm": true,
	".m4v":  true,
	".mpeg": true,
	".mpg":  true,
	".3gp":  true,

	// Fonts
	".ttf":   true,
	".otf":   true,
	".woff":  true,
	".woff2": true,
	".eot":   true,
	".fon":   true,

	// Compiled/bytecode
	".class":   true,
	".jar":     true,
	".war":     true,
	".ear":     true,
	".aar":     true,
	".ipa":     true,
	".app":     true,
	".sys":     true,
	".drv":     true,
	".ko":      true,
	".vxd":     true,
	".ocx":     true,
	".pyc":     true,
	".pyo":     true,
	".pyd":     true,
	".elc":     true,
	".hi":      true,
	".o":       true,
	".obj":     true,
	".ilk":     true,
	".pdb":     true,
	".idb":     true,
	".exp":     true,
	".lib":     true,
	".a":       true,
	".la":      true,
	".lo":      true,
	".mod":     true,
	".symvers": true,
	".order":   true,
	".build":   true,
	// Lock files (often auto-generated)
	".lock": true,

	// Node.js specific
	".node": true,

	// Office documents
	".pages":   true,
	".numbers": true,
	".key":     true,
	".keynote": true,
	".indd":    true,
	".ai":      true,
	".eps":     true,
	".sketch":  true,
	".fig":     true,
	".xd":      true,

	// Game and multimedia files
	".unity":        true,
	".unitypackage": true,
	".uasset":       true,
	".umap":         true,
	".blend":        true,
	".blend1":       true,
	".ma":           true,
	".mb":           true,
	".max":          true,
	".c4d":          true,
	".fbx":          true,
	".3ds":          true,
	".stl":          true,
	".dae":          true,
	".gltf":         true,
	".glb":          true,
	".ply":          true,
	".wrl":          true,
	".x3d":          true,
	".usd":          true,
	".usda":         true,
	".usdc":         true,
	".usdz":         true,

	// System and virtual machine files
	".vmdk":  true,
	".vhd":   true,
	".vhdx":  true,
	".qcow2": true,
	".vdi":   true,
	".ova":   true,
	".ovf":   true,
	".nvram": true,
	".vmsn":  true,
	".vmem":  true,
	".vmss":  true,
	".vswp":  true,
	".vmx":   true,
	".vmxf":  true,
	".pvm":   true,
	".hdd":   true,
	".sub":   true,

	// Database files
	".mdf":  true,
	".ldf":  true,
	".ndf":  true,
	".bak":  true,
	".trn":  true,
	".myd":  true,
	".myi":  true,
	".db3":  true,
	".sdb":  true,
	".s3db": true,
	".sl3":  true,

	// Email and communication files
	".pst":   true,
	".ost":   true,
	".msg":   true,
	".eml":   true,
	".emlx":  true,
	".mbx":   true,
	".mbox":  true,
	".vcf":   true,
	".vcard": true,

	// Backup and temporary files
	".tmp":        true,
	".temp":       true,
	".swp":        true,
	".swo":        true,
	".backup":     true,
	".old":        true,
	".save":       true,
	".sav":        true,
	".cache":      true,
	".part":       true,
	".crdownload": true,
	".download":   true,
	".partial":    true,
	".aria2":      true,

	// Security and certificate files
	".p12":        true,
	".pfx":        true,
	".pem":        true,
	".crt":        true,
	".cer":        true,
	".der":        true,
	".jks":        true,
	".keystore":   true,
	".truststore": true,
	".gpg":        true,
	".pgp":        true,
	".asc":        true,
	".sig":        true,

	// Other binary formats
	".hex":      true,
	".rom":      true,
	".img":      true,
	".toast":    true,
	".vcd":      true,
	".nrg":      true,
	".mds":      true,
	".ccd":      true,
	".cue":      true,
	".ecm":      true,
	".ape":      true,
	".tta":      true,
	".dts":      true,
	".thd":      true,
	".mlp":      true,
	".amr":      true,
	".3ga":      true,
	".mxmf":     true,
	".imy":      true,
	".awb":      true,
	".qcp":      true,
	".dvf":      true,
	".msv":      true,
	".ra":       true,
	".rm":       true,
	".ram":      true,
	".smi":      true,
	".smil":     true,
	".xap":      true,
	".jad":      true,
	".sis":      true,
	".sisx":     true,
	".prc":      true,
	".azw":      true,
	".azw3":     true,
	".kfx":      true,
	".lit":      true,
	".fb2":      true,
	".fb2.zip":  true,
	".lrf":      true,
	".lrx":      true,
	".baf":      true,
	".zno":      true,
	".tr2":      true,
	".tr3":      true,
	".pdc":      true,
	".xeb":      true,
	".ceb":      true,
	".pkg":      true,
	".xpi":      true,
	".crx":      true,
	".oex":      true,
	".nex":      true,
	".mxp":      true,
	".air":      true,
	".widget":   true,
	".wgt":      true,
	".dmp":      true,
	".hprof":    true,
	".heapdump": true,
	".core":     true,
	".minidump": true,
	".mdmp":     true,
	".wer":      true,
	".etl":      true,
	".evtx":     true,
	".evt":      true,
}

// FilenameLanguages maps specific filenames (without extension) to languages
var FilenameLanguages = func() map[string]*Language {
	const capacity = 40
	m := make(map[string]*Language, capacity)

	m["Makefile"] = &Language{
		Name:              "Makefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["makefile"] = &Language{
		Name:              "Makefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["GNUmakefile"] = &Language{
		Name:              "Makefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["Dockerfile"] = &Language{
		Name:              "Dockerfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["dockerfile"] = &Language{
		Name:              "Dockerfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["LICENSE"] = &Language{
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["LICENSE.txt"] = &Language{
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["LICENSE.md"] = &Language{
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["LICENCE"] = &Language{
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["COPYING"] = &Language{
		Name:              "License",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["README"] = &Language{
		Name:              "Readme",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["README.txt"] = &Language{
		Name:              "Readme",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["Vagrantfile"] = &Language{
		Name:              "Vagrantfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
	}
	m["Gemfile"] = &Language{
		Name:              "Gemfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
	}
	m["Rakefile"] = &Language{
		Name:              "Rakefile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "=begin",
		MultiLineEnd:      "=end",
	}
	m["Procfile"] = &Language{
		Name:              "Procfile",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["CMakeLists.txt"] = &Language{
		Name:              "CMake",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["Jenkinsfile"] = &Language{
		Name:              "Jenkinsfile",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m["CHANGELOG"] = &Language{
		Name:              "Changelog",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["CHANGELOG.md"] = &Language{
		Name:              "Changelog",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["AUTHORS"] = &Language{
		Name:              "Authors",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["CONTRIBUTORS"] = &Language{
		Name:              "Contributors",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["build.gradle"] = &Language{
		Name:              "Gradle",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m["build.gradle.kts"] = &Language{
		Name:              "Gradle Kotlin",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m["gradlew"] = &Language{
		Name:              "Shell Script",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["gradlew.bat"] = &Language{
		Name:              "Batch File",
		Extensions:        []string{},
		SingleLineComment: "REM",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["mvnw"] = &Language{
		Name:              "Shell Script",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["mvnw.cmd"] = &Language{
		Name:              "Batch File",
		Extensions:        []string{},
		SingleLineComment: "REM",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["pom.xml"] = &Language{
		Name:              "Maven POM",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "<!--",
		MultiLineEnd:      "-->",
	}
	m["composer.json"] = &Language{
		Name:              "Composer",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["package.json"] = &Language{
		Name:              "NPM Package",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["tsconfig.json"] = &Language{
		Name:              "TypeScript Config",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}
	m["webpack.config.js"] = &Language{
		Name:              "Webpack Config",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m["gruntfile.js"] = &Language{
		Name:              "Grunt",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}
	m["gulpfile.js"] = &Language{
		Name:              "Gulp",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}

	return m
}()

// HiddenFileLanguages maps hidden config files to languages
var HiddenFileLanguages = func() map[string]*Language {
	const estimatedCapacity = 20

	m := make(map[string]*Language, estimatedCapacity)
	m[".gitignore"] = &Language{
		Name:              "Git Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".dockerignore"] = &Language{
		Name:              "Docker Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".gitattributes"] = &Language{
		Name:              "Git Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".editorconfig"] = &Language{
		Name:              "EditorConfig",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".eslintrc"] = &Language{
		Name:              "ESLint Config",
		Extensions:        []string{},
		SingleLineComment: "//",
		MultiLineStart:    "/*",
		MultiLineEnd:      "*/",
	}

	m[".prettierrc"] = &Language{
		Name:              "Prettier Config",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".babelrc"] = &Language{
		Name:              "Babel Config",
		Extensions:        []string{},
		SingleLineComment: "",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".npmrc"] = &Language{
		Name:              "NPM Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".yarnrc"] = &Language{
		Name:              "Yarn Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".env"] = &Language{
		Name:              "Environment",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".env.example"] = &Language{
		Name:              "Environment",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".env.local"] = &Language{
		Name:              "Environment",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".htaccess"] = &Language{
		Name:              "Apache Config",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	m[".travis.yml"] = &Language{
		Name:              "Travis CI",
		Extensions:        []string{},
		SingleLineComment: "#",
		MultiLineStart:    "",
		MultiLineEnd:      "",
	}

	return m
}()

// GetLanguage returns the language definition for a given file extension
func GetLanguage(ext string) *Language {
	if lang, ok := languageSync.Load(ext); ok {
		return lang.(*Language)
	}
	return nil
}

// GetLanguageByFilename returns the language definition for a specific filename
func GetLanguageByFilename(filename string) *Language {
	// Check exact filename match first
	if lang, ok := fileNameLanguageSync.Load(filename); ok {
		return lang.(*Language)
	}
	// Check hidden file languages
	if lang, ok := hiddenFileLanguageSync.Load(filename); ok {
		return lang.(*Language)
	}
	return nil
}

// IsBinaryExtension checks if the file extension is a binary file
func IsBinaryExtension(ext string) bool {
	return BinaryExtensions[ext]
}
