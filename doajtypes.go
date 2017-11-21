// Generated with Chidley, https://github.com/gnewton/chidley

package main

type DOAJRecords struct {
	DOAJRecords []*DOAJRecord `xml:" record,omitempty" json:"record,omitempty"`
}

type DOAJRecord struct {
	DOAJAbstract          *DOAJAbstract          `xml:" abstract,omitempty" json:"abstract,omitempty"`
	DOAJAffiliationsList  *DOAJAffiliationsList  `xml:" affiliationsList,omitempty" json:"affiliationsList,omitempty"`
	DOAJAuthors           *DOAJAuthors           `xml:" authors,omitempty" json:"authors,omitempty"`
	DOAJDocumentType      *DOAJDocumentType      `xml:" documentType,omitempty" json:"documentType,omitempty"`
	DOAJDoi               *DOAJDoi               `xml:" doi,omitempty" json:"doi,omitempty"`
	DOAJEndPage           *DOAJEndPage           `xml:" endPage,omitempty" json:"endPage,omitempty"`
	DOAJFullTextUrl       *DOAJFullTextUrl       `xml:" fullTextUrl,omitempty" json:"fullTextUrl,omitempty"`
	DOAJIssn              *DOAJIssn              `xml:" issn,omitempty" json:"issn,omitempty"`
	DOAJIssue             *DOAJIssue             `xml:" issue,omitempty" json:"issue,omitempty"`
	DOAJJournalTitle      *DOAJJournalTitle      `xml:" journalTitle,omitempty" json:"journalTitle,omitempty"`
	DOAJKeywords          *DOAJKeywords          `xml:" keywords,omitempty" json:"keywords,omitempty"`
	DOAJLanguage          *DOAJLanguage          `xml:" language,omitempty" json:"language,omitempty"`
	DOAJPublicationDate   *DOAJPublicationDate   `xml:" publicationDate,omitempty" json:"publicationDate,omitempty"`
	DOAJPublisher         *DOAJPublisher         `xml:" publisher,omitempty" json:"publisher,omitempty"`
	DOAJPublisherRecordId *DOAJPublisherRecordId `xml:" publisherRecordId,omitempty" json:"publisherRecordId,omitempty"`
	DOAJStartPage         *DOAJStartPage         `xml:" startPage,omitempty" json:"startPage,omitempty"`
	DOAJTitle             *DOAJTitle             `xml:" title,omitempty" json:"title,omitempty"`
	DOAJVolume            *DOAJVolume            `xml:" volume,omitempty" json:"volume,omitempty"`
}

type DOAJLanguage struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJPublisher struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJJournalTitle struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJIssn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJPublicationDate struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJVolume struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJIssue struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJStartPage struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJEndPage struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJDoi struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJPublisherRecordId struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type DOAJDocumentType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJTitle struct {
	AttrLanguage string `xml:" language,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type DOAJAuthors struct {
	DOAJAuthor []*DOAJAuthor `xml:" author,omitempty" json:"author,omitempty"`
}

type DOAJAuthor struct {
	DOAJAffiliationId *DOAJAffiliationId `xml:" affiliationId,omitempty" json:"affiliationId,omitempty"`
	DOAJEmail         *DOAJEmail         `xml:" email,omitempty" json:"email,omitempty"`
	DOAJName          *DOAJName          `xml:" name,omitempty" json:"name,omitempty"`
}

type DOAJName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJEmail struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type DOAJAffiliationId struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

type DOAJAffiliationsList struct {
	DOAJAffiliationName []*DOAJAffiliationName `xml:" affiliationName,omitempty" json:"affiliationName,omitempty"`
}

type DOAJAffiliationName struct {
	AttrAffiliationId int8   `xml:" affiliationId,attr"  json:",omitempty"`
	Text              string `xml:",chardata" json:",omitempty"`
}

type DOAJAbstract struct {
	AttrLanguage string `xml:" language,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

type DOAJFullTextUrl struct {
	AttrFormat string `xml:" format,attr"  json:",omitempty"`
	Text       string `xml:",chardata" json:",omitempty"`
}

type DOAJKeywords struct {
	AttrLanguage string         `xml:" language,attr"  json:",omitempty"`
	DOAJKeyword  []*DOAJKeyword `xml:" keyword,omitempty" json:"keyword,omitempty"`
}

type DOAJKeyword struct {
	Text string `xml:",chardata" json:",omitempty"`
}
