package helpers

//* https://docs.github.com/en/rest/pulls?apiVersion=2022-11-28
//! curl -L -H "Authorization: Bearer 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" https://api.github.com/repos/tinshade/codeloom/pulls

const GITHUB_BASE_URL string = "https://api.github.com"
const FIXTURES_BASE_PATH string = "../../internal/models/fixtures"