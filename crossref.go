// documentation http://data.crossref.org/reports/help/schema_doc/4.4.1/index.html

package main

import (
	"log"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

type TemplateData struct {
	HeadData
	BodyData
}

type HeadData struct {
	DOIBatch       int64
	Timestamp      int64
	DepositorName  string
	DepositorEmail string
	Registrant     string
}

type BodyData struct {
	Journals []*Journal
}

type Journal struct {
	LanguageCode     string
	FullTitle        string
	AbbrevTitle      string
	ISSNs            []ISSN
	PublicationDates []PublicationDate
	Volume           string
	Issue            string
	Articles         []Article
}

type ISSN struct {
	Value string
	Type  string
}

type PublicationDate struct {
	Year  string
	Month string
	Day   string
	Type  string
}

type Article struct {
	Title            string
	Contributors     []Contributor
	PublicationDates []PublicationDate
	DOI              string
	URI              string
	FirstPage        string
	LastPage         string
}

type Contributor struct {
	GivenName   string
	Surname     string
	Affiliation string
}

func CreateTemplateData(depositorName, depositorEmail, registrant string, mapping map[string]PrefixAndAbbreviation, records *DOAJRecords) *TemplateData {

	templateData := new(TemplateData)

	templateData.HeadData = HeadData{
		DOIBatch:       time.Now().UTC().Unix(),
		Timestamp:      time.Now().UTC().UnixNano(),
		DepositorName:  depositorName,
		DepositorEmail: depositorEmail,
		Registrant:     registrant,
	}

	for _, record := range records.DOAJRecords {
		journal := GetOrCreateJournal(&templateData.BodyData, mapping, record)
		journal.AddArticle(mapping, record)
	}

	return templateData
}

func (j *Journal) AddArticle(mapping map[string]PrefixAndAbbreviation, record *DOAJRecord) {

	fulltextUrl, err := url.Parse(record.DOAJFullTextUrl.Text)
	if err != nil {
		log.Fatalln("Unable to parse full text url", err)
	}

	doi := mapping[j.FullTitle].Prefix + path.Base(fulltextUrl.Path)

	j.Articles = append(j.Articles, Article{
		Title:            record.DOAJTitle.Text,
		URI:              record.DOAJFullTextUrl.Text,
		FirstPage:        record.DOAJStartPage.Text,
		LastPage:         record.DOAJEndPage.Text,
		DOI:              doi,
		PublicationDates: j.PublicationDates,
		Contributors:     CreateContributors(record),
	})
}

func CreateContributors(record *DOAJRecord) []Contributor {

	idToAffiliation := make(map[int8]string)
	contributors := []Contributor{}

	for _, affiliation := range record.DOAJAffiliationsList.DOAJAffiliationName {
		idToAffiliation[affiliation.AttrAffiliationId] = strings.TrimSpace(affiliation.Text)
	}

	for _, contributor := range record.DOAJAuthors.DOAJAuthor {

		log.Println(contributor.DOAJName.Text)

		var c Contributor

		if len(strings.SplitN(contributor.DOAJName.Text, " ", 2)) == 1 {
			c = Contributor{
				Surname: contributor.DOAJName.Text,
			}
		} else {
			c = Contributor{
				GivenName: strings.SplitN(contributor.DOAJName.Text, " ", 2)[0],
				Surname:   strings.SplitN(contributor.DOAJName.Text, " ", 2)[1],
			}
		}

		if contributor.DOAJAffiliationId != nil {
			c.Affiliation = idToAffiliation[contributor.DOAJAffiliationId.Text]
		}

		contributors = append(contributors, c)
	}

	return contributors

}

func GetOrCreateJournal(bodyData *BodyData, mapping map[string]PrefixAndAbbreviation, record *DOAJRecord) *Journal {

	for i := range bodyData.Journals {
		journal := bodyData.Journals[i]
		if journal.FullTitle == record.DOAJJournalTitle.Text &&
			journal.Volume == record.DOAJVolume.Text &&
			journal.Issue == record.DOAJIssue.Text {

			return journal
		}
	}

	journal := &Journal{
		LanguageCode:     ISO6392toISO6391(record.DOAJLanguage.Text),
		FullTitle:        record.DOAJJournalTitle.Text,
		AbbrevTitle:      mapping[record.DOAJJournalTitle.Text].Abbreviation,
		ISSNs:            CreateISSNs(record),
		PublicationDates: CreatePublicationDates(record),
		Volume:           record.DOAJVolume.Text,
		Issue:            record.DOAJIssue.Text,
		Articles:         []Article{},
	}

	bodyData.Journals = append(bodyData.Journals, journal)

	return journal
}

func CreateISSNs(record *DOAJRecord) []ISSN {
	return []ISSN{
		ISSN{record.DOAJIssn.Text, "electronic"},
	}
}

func CreatePublicationDates(record *DOAJRecord) []PublicationDate {

	t, err := time.Parse("2006-01-02", record.DOAJPublicationDate.Text)
	if err != nil {
		log.Fatalln("Unable to process date", err)
	}

	return []PublicationDate{
		PublicationDate{strconv.Itoa(t.Year()), strconv.Itoa(int(t.Month())), strconv.Itoa(t.Day()), "electronic"},
	}
}

func ISO6392toISO6391(code string) string {
	switch code {
	case "aar":
		return "aa"
	case "afr":
		return "af"
	case "aka":
		return "ak"
	case "alb":
		return "sq"
	case "amh":
		return "am"
	case "ara":
		return "ar"
	case "arg":
		return "an"
	case "arm":
		return "hy"
	case "asm":
		return "as"
	case "ava":
		return "av"
	case "ave":
		return "ae"
	case "aym":
		return "ay"
	case "aze":
		return "az"
	case "bak":
		return "ba"
	case "bam":
		return "bm"
	case "baq":
		return "eu"
	case "bel":
		return "be"
	case "ben":
		return "bn"
	case "bih":
		return "bh"
	case "bis":
		return "bi"
	case "bod":
		return "bo"
	case "bos":
		return "bs"
	case "bre":
		return "br"
	case "bul":
		return "bg"
	case "bur":
		return "my"
	case "cat":
		return "ca"
	case "ces":
		return "cs"
	case "cha":
		return "ch"
	case "che":
		return "ce"
	case "chi":
		return "zh"
	case "chu":
		return "cu"
	case "chv":
		return "cv"
	case "cor":
		return "kw"
	case "cos":
		return "co"
	case "cre":
		return "cr"
	case "cym":
		return "cy"
	case "cze":
		return "cs"
	case "dan":
		return "da"
	case "deu":
		return "de"
	case "div":
		return "dv"
	case "dut":
		return "nl"
	case "dzo":
		return "dz"
	case "ell":
		return "el"
	case "eng":
		return "en"
	case "epo":
		return "eo"
	case "est":
		return "et"
	case "eus":
		return "eu"
	case "ewe":
		return "ee"
	case "fao":
		return "fo"
	case "fas":
		return "fa"
	case "fij":
		return "fj"
	case "fin":
		return "fi"
	case "fra":
		return "fr"
	case "fre":
		return "fr"
	case "fry":
		return "fy"
	case "ful":
		return "ff"
	case "geo":
		return "ka"
	case "ger":
		return "de"
	case "gla":
		return "gd"
	case "gle":
		return "ga"
	case "glg":
		return "gl"
	case "glv":
		return "gv"
	case "gre":
		return "el"
	case "grn":
		return "gn"
	case "guj":
		return "gu"
	case "hat":
		return "ht"
	case "hau":
		return "ha"
	case "heb":
		return "he"
	case "her":
		return "hz"
	case "hin":
		return "hi"
	case "hmo":
		return "ho"
	case "hrv":
		return "hr"
	case "hun":
		return "hu"
	case "hye":
		return "hy"
	case "ibo":
		return "ig"
	case "ice":
		return "is"
	case "ido":
		return "io"
	case "iii":
		return "ii"
	case "iku":
		return "iu"
	case "ile":
		return "ie"
	case "ina":
		return "ia"
	case "ind":
		return "id"
	case "ipk":
		return "ik"
	case "isl":
		return "is"
	case "ita":
		return "it"
	case "jav":
		return "jv"
	case "jpn":
		return "ja"
	case "kal":
		return "kl"
	case "kan":
		return "kn"
	case "kas":
		return "ks"
	case "kat":
		return "ka"
	case "kau":
		return "kr"
	case "kaz":
		return "kk"
	case "khm":
		return "km"
	case "kik":
		return "ki"
	case "kin":
		return "rw"
	case "kir":
		return "ky"
	case "kom":
		return "kv"
	case "kon":
		return "kg"
	case "kor":
		return "ko"
	case "kua":
		return "kj"
	case "kur":
		return "ku"
	case "lao":
		return "lo"
	case "lat":
		return "la"
	case "lav":
		return "lv"
	case "lim":
		return "li"
	case "lin":
		return "ln"
	case "lit":
		return "lt"
	case "ltz":
		return "lb"
	case "lub":
		return "lu"
	case "lug":
		return "lg"
	case "mac":
		return "mk"
	case "mah":
		return "mh"
	case "mal":
		return "ml"
	case "mao":
		return "mi"
	case "mar":
		return "mr"
	case "may":
		return "ms"
	case "mkd":
		return "mk"
	case "mlg":
		return "mg"
	case "mlt":
		return "mt"
	case "mon":
		return "mn"
	case "mri":
		return "mi"
	case "msa":
		return "ms"
	case "mya":
		return "my"
	case "nau":
		return "na"
	case "nav":
		return "nv"
	case "nbl":
		return "nr"
	case "nde":
		return "nd"
	case "ndo":
		return "ng"
	case "nep":
		return "ne"
	case "nld":
		return "nl"
	case "nno":
		return "nn"
	case "nob":
		return "nb"
	case "nor":
		return "no"
	case "nya":
		return "ny"
	case "oci":
		return "oc"
	case "oji":
		return "oj"
	case "ori":
		return "or"
	case "orm":
		return "om"
	case "oss":
		return "os"
	case "pan":
		return "pa"
	case "per":
		return "fa"
	case "pli":
		return "pi"
	case "pol":
		return "pl"
	case "por":
		return "pt"
	case "pus":
		return "ps"
	case "que":
		return "qu"
	case "roh":
		return "rm"
	case "ron":
		return "ro"
	case "rum":
		return "ro"
	case "run":
		return "rn"
	case "rus":
		return "ru"
	case "sag":
		return "sg"
	case "san":
		return "sa"
	case "sin":
		return "si"
	case "slk":
		return "sk"
	case "slo":
		return "sk"
	case "slv":
		return "sl"
	case "sme":
		return "se"
	case "smo":
		return "sm"
	case "sna":
		return "sn"
	case "snd":
		return "sd"
	case "som":
		return "so"
	case "sot":
		return "st"
	case "spa":
		return "es"
	case "sqi":
		return "sq"
	case "srd":
		return "sc"
	case "srp":
		return "sr"
	case "ssw":
		return "ss"
	case "sun":
		return "su"
	case "swa":
		return "sw"
	case "swe":
		return "sv"
	case "tah":
		return "ty"
	case "tam":
		return "ta"
	case "tat":
		return "tt"
	case "tel":
		return "te"
	case "tgk":
		return "tg"
	case "tgl":
		return "tl"
	case "tha":
		return "th"
	case "tib":
		return "bo"
	case "tir":
		return "ti"
	case "ton":
		return "to"
	case "tsn":
		return "tn"
	case "tso":
		return "ts"
	case "tuk":
		return "tk"
	case "tur":
		return "tr"
	case "twi":
		return "tw"
	case "uig":
		return "ug"
	case "ukr":
		return "uk"
	case "urd":
		return "ur"
	case "uzb":
		return "uz"
	case "ven":
		return "ve"
	case "vie":
		return "vi"
	case "vol":
		return "vo"
	case "wel":
		return "cy"
	case "wln":
		return "wa"
	case "wol":
		return "wo"
	case "xho":
		return "xh"
	case "yid":
		return "yi"
	case "yor":
		return "yo"
	case "zha":
		return "za"
	case "zho":
		return "zh"
	case "zul":
		return "zu"
	default:
		return ""
	}
}
