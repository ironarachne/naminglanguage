package naminglanguage

var (
	articles     = []string{"a", "an", "the", "what", "why", "when", "who", "how"}
	adverbs      = []string{"no", "yes", "also"}
	adjectives   = []string{"fortunate", "industrious", "mighty", "beloved", "compassionate", "earnest", "noble", "bitter", "constant", "free", "northern", "southern", "blessed", "desired", "famous", "hallowed", "loyal", "brave", "divine", "happy", "peaceful", "powerful", "ready", "wealthy", "praiseworthy", "sharp", "swift", "wise", "prayerful", "shining", "valiant", "protecting", "small", "big", "victorious", "worthy", "pure"}
	nouns        = []string{"arrow", "crown", "father", "giver", "healer", "laurel", "man", "rose", "spear", "sword", "battle", "defender", "fighter", "helper", "leader", "pearl", "ruler", "staff", "traveler", "bearer", "dweller", "forest", "guardian", "home", "lily", "person", "runner", "steward", "twin", "brightness", "earth", "gate", "hammer", "horse", "lover", "protector", "smith", "stranger", "warrior", "counselor", "farmer", "gift", "harvester", "keeper", "rock", "son", "stronghold", "wolf", "day", "night"}
	verbs        = []string{"heal", "kill", "borrow", "steal", "give", "speak", "make", "is", "be", "have", "know", "use", "think", "like"}
	conjunctions = []string{"and", "as", "of", "but", "in"}
	pronouns     = []string{"him", "her", "they", "we", "she", "he"}
)

var ( //Geography Terms
	ecosystems     = []string{"plain", "savannah", "meadow", "field", "prairie", "steppe"}
	forests        = []string{"forest", "woods", "grove", "copse", "stand"}
	hillsMountains = []string{"hill", "down", "mountain", "peak", "range", "plateau", "crest", "mesa"}
	valleys        = []string{"valley", "canyon", "cliff"}
	harshTerrain   = []string{"marsh", "swamp", "desert", "wilderness", "jungle"}
	rivers         = []string{"river", "stream", "brook", "estuary", "channel", "rapids", "portage", "source", "confluence", "delta", "bank", "waterfall", "spring"}
	seasNLakes     = []string{"ocean", "sea", "bay", "harbor", "lagoon", "gulf", "straight", "lake", "pond"}
	coastlines     = []string{"coast", "beach", "peninsula", "cape", "point"}
	islands        = []string{"island", "archipelago", "atoll"}
	localFeatures  = []string{"rock", "tree", "bridge", "ford", "dam", "park", "oasis"}
)

var ( //Human terms
	settlements = []string{"city", "town", "village", "hamlet", "colony", "port", "market", "capital", "kingdom", "empire", "nation"}
	military    = []string{"fort", "castle", "wall", "camp", "tower"}
	religion    = []string{"shrine", "oracle", "temple", "church", "chapel", "monastary"}
	other       = []string{"mine", "inn", "stopping point", "post", "lodge", "mill", "house", "manor", "estate", "crossing", "farm", "orchard"}
	roads       = []string{"road", "highway", "trail", "way", "path", "causeway", "bridge"}
)

var ( //Adjectives
	age          = []string{"new", "old", "ancient"}
	sizeShape    = []string{"round", "flat", "wide", "narrow", "small", "large", "split", "vast", "greater", "lesser"}
	location     = []string{"high", "low", "central", "upper", "lower", "near", "far"}
	directions   = []string{"north", "south", "east", "west"}
	greatness    = []string{"", "great", "grand", "glorious", "noble", "holy", "royal"}
	descriptive  = []string{"windy", "sandy", "cloudy", "foggy", "rainy", "snowy", "dry", "wet", "good", "twisted", "twisting", "winding", "fertile", "beautiful", "pleasant", "quiet", "gold", "golden", "silver", "jeweled", "jewel"}
	colors       = []string{"white", "black", "red", "brown", "orange", "yellow", "green", "blue", "purple", "gold", "silver", "grey", "pink", "tan"}
	numbers      = []string{"full", "half", "quarter", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	waterQuality = []string{"fast", "slow", "clear", "muddy"}
)

var ( //Common Nouns
	animals            = []string{"horse", "ox", "chicken", "sheep", "goat", "lion", "dog", "fox", "deer", "hawk", "eagle", "dragon", "fist", "turtle"}
	substance          = []string{"earth", "water", "ice", "stone", "fire", "ashes", "smoke", "dust", "salt", "sand"}
	professions        = []string{"king", "emperor", "lord", "knight", "warrior", "soldier", "priest", "holy man", "hunter", "fisherman", "sailor", "baker", "forager"}
	typesOfTrees       = []string{"aspen", "birch", "cedar", "elm", "oak", "palm", "pine", "willow"}
	flowers            = []string{"rose", "lily", "tulip", "sunflower", "lilac", "lavender", "petunia"}
	gems               = []string{"diamond", "emerald", "ruby", "sapphire", "topaz", "opal", "amethyst", "agate", "lapis lazuli", "amber", "quartz", "onyx", "pearl", "garnet"}
	positiveAttributes = []string{"grace", "hope", "beauty", " strength", "courage"}
	spaceterms         = []string{"planet", "moon", "star", "sun", "asteroid", "asteroid belt", "asteroid field", "comet", "system", "galaxy", "nebula", "ship", "station"}
)
