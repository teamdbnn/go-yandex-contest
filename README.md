# go-yandex-contest
yandex contest api client (built from swagger https://api.contest.yandex.net/api/public/swagger-ui.html)



## Methods
- [x] GET /competitions/{competitionId}/participants Get registered participants of competition
- [ ] POST /competitions/{competitionId}/participants Register participants into competition
- [x] GET /contests/{contestId} Get information about contest by id
- [ ] GET /contests/{contestId}/clarifications Get clarifications in contest by contest id
- [ ] GET /contests/{contestId}/messages Get messages in contest by contest id
- [ ] POST /contests/{contestId}/messages Post new question to contest
- [x] GET /contests/{contestId}/participants Get contest participants
- [ ] POST /contests/{contestId}/participants Register for contest
- [ ] GET /contests/{contestId}/participants/{participantId} Get information about participant by id
- [ ] PUT /contests/{contestId}/participants/{participantId} Start the contest for participant
- [ ] DELETE /contests/{contestId}/participants/{participantId} Unregister user from contest by id
- [ ] PATCH /contests/{contestId}/participants/{participantId} Update participant in contest
- [ ] GET /contests/{contestId}/participants/{participantId}/stats Get participant stats
- [ ] GET /contests/{contestId}/participation Get information about your status
- [ ] PUT /contests/{contestId}/participation Start the contest
- [ ] DELETE /contests/{contestId}/participation Unregister yourself from contest
- [ ] GET /contests/{contestId}/problems Get contest's problems by contest id
- [ ] GET /contests/{contestId}/standings Get contest's standings by contest id
- [ ] GET /contests/{contestId}/standings/my Get user's standings by contest id
- [ ] GET /contests/{contestId}/submissions Get your submissions to contest by contest id
- [ ] POST /contests/{contestId}/submissions Send a submission to contest
- [ ] GET /contests/{contestId}/submissions/{runId} Get detailed information about submission
- [ ] GET /contests/{contestId}/submissions/{runId}/{testName}/answer Get full answer file in contest by test name
- [ ] GET /contests/{contestId}/submissions/{runId}/{testName}/input Get full input file in contest by test name
- [ ] GET /contests/{contestId}/submissions/{runId}/{testName}/output Get output file in contest by test name
- [ ] GET /contests/{contestId}/submissions/{runId}/full Get full information about submission
- [ ] GET /contests/{contestId}/submissions/{runId}/source Get source code of submission
- [ ] HEAD /contests/{contestId}/submissions/{runId}/source Get source code of submission
- [ ] POST /contests/{contestId}/submissions/lazy Send a submission to contest lazy
- [ ] GET /contests/{contestId}/submissions/multiple Get full information about multiple submissions
- [ ] GET /contests/neurips/{contestId}/submissions/all Get all contest submissions report
- [ ] GET /contests/neurips/{contestId}/submissions/my Get your submissions to contest by contest id
- [ ] GET /participants/{participantId}/stats Get participant stats
- [ ] GET /service/capacity Get contest submissions queue capacity
- [x] POST /submissions/{submissionId}/rejudge Rejudge submission
- [x] POST /user/{userId}/generate-password Generate new password for internal user
