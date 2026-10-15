// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package xtwitterscraper_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/Xquik-dev/x-twitter-scraper-go"
)

func TestXJobGetFixture(t *testing.T) {
	client, request := fixtureClient(t, "getJob.json")
	res, err := client.X.Jobs.Get(context.Background(), "1899202395210649861", xtwitterscraper.XJobGetParams{
		Fields: xtwitterscraper.String("id,text,author.username"),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/jobs/1899202395210649861", "fields=id%2Ctext%2Cauthor.username&omit=entities%2Cmedia")
	expectDecoded(t, res)
	if res.Job.ID != "1899202395210649861" ||
		res.Job.Title != "Engineer" ||
		res.Job.URL != "https://x.com/i/jobs/1899202395210649861" ||
		res.Job.ApplyURL != "https://xquik.com/example" ||
		res.Job.Location != "United States" ||
		res.Job.LocationType != "remote" ||
		res.Job.EmploymentType != "full_time" ||
		res.Job.SeniorityLevel != "lead" ||
		res.Job.JobFunction != "software_engineering" ||
		res.Job.FormattedSalary != "Starting at $225K" ||
		res.Job.SalaryMin != 225000 ||
		res.Job.SalaryCurrencyCode != "USD" ||
		res.Job.SalaryInterval != "annually" ||
		!res.Job.Featured ||
		res.Job.ShortDescription != "Commonware is hiring founding engineers." ||
		res.Job.Description != "What is Commonware? We build open primitives." ||
		res.Job.DescriptionMarkdown != "# What is Commonware?" ||
		res.Job.Account.ID != "1812126202804711425" ||
		res.Job.Account.Username != "commonwarexyz" ||
		res.Job.Account.Name != "commonware" ||
		!res.Job.Account.Verified ||
		res.Job.Account.VerifiedType != "Business" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXJobListLocationsFixture(t *testing.T) {
	client, request := fixtureClient(t, "getJobLocations.json")
	res, err := client.X.Jobs.ListLocations(context.Background(), xtwitterscraper.XJobListLocationsParams{
		Q:      "new york",
		Fields: xtwitterscraper.String("id,text,author.username"),
		Limit:  xtwitterscraper.Int(5),
		Omit:   xtwitterscraper.String("entities,media"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/jobs/locations", "fields=id%2Ctext%2Cauthor.username&limit=5&omit=entities%2Cmedia&q=new+york")
	expectDecoded(t, res)
	if len(res.Locations) != 2 ||
		res.Count != 2 ||
		res.Locations[0].ID != "1720462281849397612" ||
		res.Locations[0].Name != "New York, US" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}

func TestXJobSearchFixture(t *testing.T) {
	client, request := fixtureClient(t, "searchJobs.json")
	res, err := client.X.Jobs.Search(context.Background(), xtwitterscraper.XJobSearchParams{
		Q:              "engineer",
		Company:        xtwitterscraper.String("SpaceX"),
		Cursor:         xtwitterscraper.String("cursor"),
		EmploymentType: xtwitterscraper.String("full_time"),
		Fields:         xtwitterscraper.String("id,text,author.username"),
		Location:       xtwitterscraper.String("Berlin"),
		LocationID:     xtwitterscraper.String("1720462281849397612"),
		LocationType:   xtwitterscraper.String("remote,hybrid"),
		Omit:           xtwitterscraper.String("entities,media"),
		PageSize:       xtwitterscraper.Int(10),
		Seniority:      xtwitterscraper.String("senior,lead"),
	})
	if err != nil {
		t.Fatal(err)
	}
	request.expect(t, http.MethodGet, "/api/v1/x/jobs/search", "company=SpaceX&cursor=cursor&employmentType=full_time&fields=id%2Ctext%2Cauthor.username&location=Berlin&locationId=1720462281849397612&locationType=remote%2Chybrid&omit=entities%2Cmedia&pageSize=10&q=engineer&seniority=senior%2Clead")
	expectDecoded(t, res)
	if len(res.Jobs) != 1 ||
		!res.HasNextPage ||
		res.NextCursor != "DAACCgACGRElMJcAAA" ||
		res.Jobs[0].ID != "1899202395210649861" ||
		res.Jobs[0].Title != "Engineer" ||
		res.Jobs[0].URL != "https://x.com/i/jobs/1899202395210649861" ||
		res.Jobs[0].ApplyURL != "https://xquik.com/example" ||
		res.Jobs[0].Location != "United States" ||
		res.Jobs[0].FormattedSalary != "Starting at $225K" ||
		res.Jobs[0].SalaryMin != 225000 ||
		res.Jobs[0].SalaryCurrencyCode != "USD" ||
		res.Jobs[0].SalaryInterval != "annually" ||
		res.Jobs[0].Account.ID != "1812126202804711425" ||
		res.Jobs[0].Account.Username != "commonwarexyz" ||
		res.Jobs[0].Account.Name != "commonware" ||
		!res.Jobs[0].Account.Verified ||
		res.Jobs[0].Account.VerifiedType != "Business" {
		t.Fatalf("decoded %s", res.RawJSON())
	}
}
