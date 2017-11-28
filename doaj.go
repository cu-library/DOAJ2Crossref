// Generated with Chidley, https://github.com/gnewton/chidley

package main

import (
	"encoding/xml"
	"errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// DOAJRecords is the root
type DOAJRecords struct {
	DOAJRecords []*DOAJRecord `xml:" record,omitempty" json:"record,omitempty"`
}

// DOAJRecord holds article metadata
type DOAJRecord struct {
	DOAJAbstract          *DOAJAbstract          `xml:" abstract,omitempty" json:"abstract,omitempty"`
	DOAJAffiliationsList  *DOAJAffiliationsList  `xml:" affiliationsList,omitempty" json:"affiliationsList,omitempty"`
	DOAJAuthors           *DOAJAuthors           `xml:" authors,omitempty" json:"authors,omitempty"`
	DOAJDocumentType      *DOAJDocumentType      `xml:" documentType,omitempty" json:"documentType,omitempty"`
	DOAJDoi               *DOAJDoi               `xml:" doi,omitempty" json:"doi,omitempty"`
	DOAJEndPage           *DOAJEndPage           `xml:" endPage,omitempty" json:"endPage,omitempty"`
	DOAJFullTextURL       *DOAJFullTextURL       `xml:" fullTextUrl,omitempty" json:"fullTextUrl,omitempty"`
	DOAJIssn              *DOAJIssn              `xml:" issn,omitempty" json:"issn,omitempty"`
	DOAJIssue             *DOAJIssue             `xml:" issue,omitempty" json:"issue,omitempty"`
	DOAJJournalTitle      *DOAJJournalTitle      `xml:" journalTitle,omitempty" json:"journalTitle,omitempty"`
	DOAJKeywords          *DOAJKeywords          `xml:" keywords,omitempty" json:"keywords,omitempty"`
	DOAJLanguage          *DOAJLanguage          `xml:" language,omitempty" json:"language,omitempty"`
	DOAJPublicationDate   *DOAJPublicationDate   `xml:" publicationDate,omitempty" json:"publicationDate,omitempty"`
	DOAJPublisher         *DOAJPublisher         `xml:" publisher,omitempty" json:"publisher,omitempty"`
	DOAJPublisherRecordID *DOAJPublisherRecordID `xml:" publisherRecordId,omitempty" json:"publisherRecordId,omitempty"`
	DOAJStartPage         *DOAJStartPage         `xml:" startPage,omitempty" json:"startPage,omitempty"`
	DOAJTitle             *DOAJTitle             `xml:" title,omitempty" json:"title,omitempty"`
	DOAJVolume            *DOAJVolume            `xml:" volume,omitempty" json:"volume,omitempty"`
}

// DOAJLanguage is the article language
type DOAJLanguage struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJPublisher is the article publisher
type DOAJPublisher struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJJournalTitle is the journal title
type DOAJJournalTitle struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJIssn is the article title
type DOAJIssn struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJPublicationDate is the journal/article publication date.
type DOAJPublicationDate struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJVolume is the journal volume
type DOAJVolume struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJIssue is the journal issue
type DOAJIssue struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJStartPage is the start page
type DOAJStartPage struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJEndPage is the end page
type DOAJEndPage struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJDoi is the doi
type DOAJDoi struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJPublisherRecordID is the publisher record id
type DOAJPublisherRecordID struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

// DOAJDocumentType is the document type
type DOAJDocumentType struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJTitle is the article title
type DOAJTitle struct {
	AttrLanguage string `xml:" language,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

// DOAJAuthors holds the authors
type DOAJAuthors struct {
	DOAJAuthor []*DOAJAuthor `xml:" author,omitempty" json:"author,omitempty"`
}

// DOAJAuthor is an article author
type DOAJAuthor struct {
	DOAJAffiliationID *DOAJAffiliationID `xml:" affiliationId,omitempty" json:"affiliationId,omitempty"`
	DOAJEmail         *DOAJEmail         `xml:" email,omitempty" json:"email,omitempty"`
	DOAJName          *DOAJName          `xml:" name,omitempty" json:"name,omitempty"`
}

// DOAJName is the author name
type DOAJName struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJEmail is the author email
type DOAJEmail struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// DOAJAffiliationID is the affiliation ID
type DOAJAffiliationID struct {
	Text int8 `xml:",chardata" json:",omitempty"`
}

// DOAJAffiliationsList holds the affiliations
type DOAJAffiliationsList struct {
	DOAJAffiliationName []*DOAJAffiliationName `xml:" affiliationName,omitempty" json:"affiliationName,omitempty"`
}

// DOAJAffiliationName maps ids to names
type DOAJAffiliationName struct {
	AttrAffiliationID int8   `xml:" affiliationId,attr"  json:",omitempty"`
	Text              string `xml:",chardata" json:",omitempty"`
}

// DOAJAbstract is the article abstract
type DOAJAbstract struct {
	AttrLanguage string `xml:" language,attr"  json:",omitempty"`
	Text         string `xml:",chardata" json:",omitempty"`
}

// DOAJFullTextURL is the full text url
type DOAJFullTextURL struct {
	AttrFormat string `xml:" format,attr"  json:",omitempty"`
	Text       string `xml:",chardata" json:",omitempty"`
}

// DOAJKeywords holds the article keywords
type DOAJKeywords struct {
	AttrLanguage string         `xml:" language,attr"  json:",omitempty"`
	DOAJKeyword  []*DOAJKeyword `xml:" keyword,omitempty" json:"keyword,omitempty"`
}

// DOAJKeyword is an article keyword
type DOAJKeyword struct {
	Text string `xml:",chardata" json:",omitempty"`
}

// LoadDOAJ loads the DOAJ data from an XML file into a DOAJRecord struct.
func LoadDOAJ(xmlFilePath string) (*DOAJRecords, error) {

	records := new(DOAJRecords)
	absoluteXMLFilePath, err := filepath.Abs(xmlFilePath)

	xmlFile, err := os.Open(absoluteXMLFilePath)
	if err != nil {
		return records, err
	}
	defer xmlFile.Close()

	xmlDecoder := xml.NewDecoder(xmlFile)
	err = xmlDecoder.Decode(records)
	if err != nil {
		return records, err
	}

	return records, nil
}

// Validate looks through a DOAJ struct to ensure no records would fail crossref validation.
func (r *DOAJRecords) Validate() bool {

	ok := true

	for _, record := range r.DOAJRecords {
		err := record.validate()
		if err != nil {
			log.Printf("\"%v\", URL: %v\n", record.DOAJTitle.Text, record.DOAJFullTextURL.Text)
			log.Println(err)
			ok = false
		}
	}

	return ok
}

func (r *DOAJRecord) validate() error {

	//Check if publication date is not empty and parse-able.
	if strings.TrimSpace(r.DOAJPublicationDate.Text) == "" {
		return errors.New("publication date is empty")
	}
	_, err := time.Parse("2006-01-02", r.DOAJPublicationDate.Text)
	if err != nil {
		return err
	}

	//Check if URL not empty and parse-able.
	if strings.TrimSpace(r.DOAJFullTextURL.Text) == "" {
		return errors.New("fulltext URL is empty")
	}
	_, err = url.Parse(r.DOAJFullTextURL.Text)
	if err != nil {
		return err
	}

	// Check if the record has no authors.
	if len(r.DOAJAuthors.DOAJAuthor) == 0 {
		return errors.New("no authors")
	}

	// Check if the record has no journal title.
	if strings.TrimSpace(r.DOAJJournalTitle.Text) == "" {
		return errors.New("journal title is empty")
	}

	// Check if the record has no volume.
	if strings.TrimSpace(r.DOAJVolume.Text) == "" {
		return errors.New("volume is empty")
	}

	// Check if the record has no issue.
	if strings.TrimSpace(r.DOAJIssue.Text) == "" {
		return errors.New("issue is empty")
	}

	return nil

}
