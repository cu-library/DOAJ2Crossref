package main

const templateSkeleton string = `<?xml version="1.0" encoding="UTF-8"?>
<doi_batch version="4.4.1" 
           xmlns="http://www.crossref.org/schema/4.4.1"
           xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
           xsi:schemaLocation="http://www.crossref.org/schema/4.4.1 http://www.crossref.org/schemas/crossref4.4.1.xsd">
	<head>
        {{- with .HeadData}}
		<doi_batch_id>{{.DOIBatch}}</doi_batch_id>
		<timestamp>{{.Timestamp}}</timestamp>
		<depositor>
			<depositor_name>{{.DepositorName}}</depositor_name>
			<email_address>{{.DepositorEmail}}</email_address>
		</depositor>
		<registrant>{{.Registrant}}</registrant>
        {{- end}}
	</head>
	<body>
		{{- range .Journals }}
		<journal>
			<journal_metadata language="{{.LanguageCode}}">
				<full_title>{{.FullTitle}}</full_title>
				{{- if .AbbrevTitle}}
				<abbrev_title>{{.AbbrevTitle}}</abbrev_title>
				{{end}}
				{{- range .ISSNs}}
				<issn media_type="{{.Type}}">{{.Value}}</issn>
				{{- end}}
			</journal_metadata>
			<journal_issue>
				{{- range .PublicationDates}}
				<publication_date media_type="{{.Type}}">
					<month>{{.Month}}</month>
					<day>{{.Day}}</day>
					<year>{{.Year}}</year>
				</publication_date>
				{{- end}}
				<journal_volume>
					<volume>{{.Volume}}</volume>
				</journal_volume>
				<issue>{{.Issue}}</issue>
			</journal_issue>
			{{- range .Articles}}
			<journal_article publication_type="full_text">
				<titles>
					<title>{{.Title}}</title>
				</titles>
				<contributors>
				{{- range .Contributors}}
					<person_name sequence="{{.Sequence}}" contributor_role="{{.Role}}">
{{- if .GivenName}}{{"\n"}}						<given_name>{{.GivenName}}</given_name>{{end}}
						<surname>{{.Surname}}</surname>
{{- if .Affiliation}}{{"\n"}}						<affiliation>{{.Affiliation}}</affiliation>{{end}}
{{- if .ORCID}}{{"\n"}}						<ORCID>{{.ORCID}}</ORCID>{{end}}
					</person_name>
				{{- end}}
				</contributors>
				{{- range .PublicationDates}}
				<publication_date media_type="{{.Type}}">
					<month>{{.Month}}</month>
					<day>{{.Day}}</day>
					<year>{{.Year}}</year>
				</publication_date>
				{{- end}}
				<pages>
					<first_page>{{.FirstPage}}</first_page>
					{{- if .LastPage}}
					<last_page>{{.LastPage}}</last_page>{{end}}
				</pages>
				<doi_data>
					<doi>{{.DOI}}</doi>
					<resource>{{.URI}}</resource>
				</doi_data>
			</journal_article>
			{{- end}}
		</journal>
		{{- end}}/
	</body>
</doi_batch>
`
