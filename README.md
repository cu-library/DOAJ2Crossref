# DOAJ2Crossref
Using XML from the TIM Review DOAJ export tool, create Crossref-ready XML (including DOIs).

```
Usage of ./DOAJ2Crossref:
  -config string
        Path to config file. (default "config.json")
  -depositor string
        Name of the organization registering the DOIs. The name placed in this element should match the name under which a depositing organization has registered with CrossRef.
  -email string
        Email address to which batch success and/or error messages are sent. It is recommended that this address be unique to a position within the organization submitting data (e.g. "doi@...") rather than unique to a person. In this way, the alias for delivery of this mail can be changed as responsibility for submission of DOI data within the organization changes from one person to another.
  -in string
        Path to DOAJ XML file. (default "DOAJ.xml")
  -out string
        Path to which the output XML file will be written. (default "crossref.xml")
  -registrant string
        The organization that owns the information being registered.
  -report string
        Path to which the report csv file will be written. (default "report.csv")

```

This tool takes an XML input file from the TIM Review DOAJ export tool and transforms it into Crossref-ready XML. 

DOIs are assigned by looking at the journal title of each record, and using a mapping to DOI prefixes in config.json.

The tool also creates a CSV report of URIs to DOIs. 

ORCIDs are added to the Crossref XML output from mappings in the config.json file. 

## Assumptions and Notes

* All publication dates are of type "online".
* Every DOAJ record only has one ISSN, of type "electronic".
* The DOI is generated like this:
    ```golang 
    doi := prefix + path.Base(fulltextURL.Path)
    ```
    
    For example:
	
    ```xml 
    <fullTextUrl format="html">http://review.ca/a/long/path/99</fullTextUrl>
    ```
	
    with prefix `10.11000/review`
    would generate this DOI: 
    `10.11000/review99`
* If the start page is empty in the input, a start page of 1 is assigned.
* Mononymous people have their name mapped to crossref surname; given_name is left empty.
* ORCIDs from config.json are always prefixed with `https://orcid.org/`
* All authors in input are given contributor_role author.
* Author sequence in output is defined by the 'physical' sequence in input. The first author element is mapped to a person_name with sequence set to first.
* Publication date of article and containing journal are always the same.

## config.json

Here's an example config.json:

```JSON
{
        "mappings": [
                {
                        "journalTitle": "A Review Journal",
                        "prefix": "10.11000/review",
                        "abbreviatedJournalTitle":"R.J."
                },
                {
                        "journalTitle": "Code Resources",
                        "prefix": "10.11001/coderesources",
                        "abbreviatedJournalTitle": "C.R."
                }
        ],
        "orcids": [
                {
                        "name": "Ada Lovelace",
                        "orcid": "0000-0000-0000-0000"
                }
        ]
}
```

The config file lets the user define how journal titles are mapped to orcids, and how orcids are mapped to authors in the output.

