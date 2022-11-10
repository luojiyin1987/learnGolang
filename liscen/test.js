const choices = {
  q1: null,
  q2a: null,
  q2b: null,
  q2c: null,
  q3: null,
  q4a: null,
  q4b: null,
  q5: null,
  q6: null,
  q7: null,
};

const qs = ["q2a", "q2b", "q2c", "q3", "q4a", "q4b", "q5", "q6", "q7"];
const stepsEnabled = [
  true,
  true,
  false,
  false,
  true,
  true,
  false,
  true,
  true,
  true,
];

const ONE_STRONG_LICENCES = "q1strong",
  TWO_NO_COPYLEFT = "q2anocopyleft",
  TWO_STRONG_COPYLEFT = "q2bstrong",
  TWO_WEAK_COPYLEFT = "q2bweak",
  TWO_WEAK_MODULE = "q2cmod",
  TWO_WEAK_FILE = "q2cfile",
  TWO_WEAK_LIB = "q2clib",
  THREE_JURISDICTION = "q3juris",
  FOUR_GRANT_PATENTS = "q4apat",
  FOUR_PATENT_RET = "q4bpatret",
  FIVE_ENHANCED_ATTR = "q5enhattr",
  SIX_NO_LOOPHOLE = "q6noloophole",
  SEVEN_NO_PROMO = "q7nopromo";

function enabledSteps(index, enabled) {
  stepsEnabled[index] = enabled;
}

const conditionsReuseQs = [
  TWO_NO_COPYLEFT,
  TWO_STRONG_COPYLEFT,
  TWO_WEAK_COPYLEFT,
  TWO_WEAK_MODULE,
  TWO_WEAK_FILE,
  TWO_WEAK_LIB,
];
const simpleYesNoQs = [
  THREE_JURISDICTION,
  FOUR_GRANT_PATENTS,
  FOUR_PATENT_RET,
  FIVE_ENHANCED_ATTR,
  SIX_NO_LOOPHOLE,
  SEVEN_NO_PROMO,
];

const limitingQs = [ONE_STRONG_LICENCES];
let loadedLicenceData = [];
let loadedLimitedLicenceData = [];



const fs = require("fs");

let readData = fs.readFileSync("./matrix.json");
//console.log(readData)
let stuOut = JSON.parse(readData);

//console.log("stuOut", stuOut.feed.entry);

function initLicences(allLicences) {
  loadedLicenceData = allLicences.feed.entry;
  //console.log("allLicences", allLicences.feed.entry)
  initScores(allLicences.feed.entry);

  //console.log("loadedLicenceData", loadedLicenceData)
  console.log("------------------");
  //console.log("scores", scores);
}

function initScores(allApplicableLicences) {
  scores = allApplicableLicences.map(function (item) {
    return {
      title: item.title,
      score: -1,
    };
  });
}


function processChoice(formFieldId, fullChoice) {
  choices[formFieldId] = fullChoice;
  console.log("choices", choices[formFieldId]);
  console.log("isLimitingQuestion(fullChoice)", isLimitingQuestion(fullChoice));

  //第一个问题选择 想过处理
  if (isLimitingQuestion(fullChoice)) prepareLicencesList(fullChoice);

  updateForm(fullChoice);

  displayLicences();
}

function isLimitingQuestion(question) {
  return limitingQs.includes(question.split("_")[0]);
}

function prepareLicencesList(fullChoice) {
  const choice = fullChoice.split("_")[1];
  console.log("loadedLicenceData prepareLicencesList", loadedLicenceData);
  console.log("choice", choice);
  if (choice != 1) initScores(loadedLicenceData);
  // limit list of licences to those matching the req.
  else
    initScores(
      loadedLimitedLicenceData.length
        ? loadedLimitedLicenceData
        : loadedLicenceData.filter(function (item) {
            return 1 == processLimitingQuestion(fullChoice, item);
          })
    );
}

function processLimitingQuestion(fullChoice, licenceData) {
  fullChoice = fullChoice.split("_");

  let newMatch = 0;
  console.log("processLimitingQuestion licenceData", licenceData, fullChoice);
  if (fullChoice[1] != 1 || licenceData.content.includes(fullChoice[0]))
    newMatch++;

  console.log("processLimitingQuestion newMatch", newMatch);
  return newMatch;
}

function updateForm(fullChoice) {
  fullChoice = fullChoice.split("_");
  console.log("updateForm, fullChoice", qs, choices, fullChoice);
  if (fullChoice[0] == TWO_NO_COPYLEFT) {
    if (fullChoice[1] == 0) {
      openBox("q2b");
      closeBox("q2c");
    } else {
      closeBox("q2b");
      closeBox("q2c");
    }
  } else if (fullChoice[0] == TWO_STRONG_COPYLEFT) closeBox("q2c");
  else if (fullChoice[0] == TWO_WEAK_COPYLEFT) openBox("q2c");
  else if (fullChoice[0] == FOUR_GRANT_PATENTS)
    if (fullChoice[1] == DONT_CARE || fullChoice[1] == 1) openBox("q4b");
    else closeBox("q4b");

  console.log("updateForm, qs choices", qs, choices);
}

function openBox(boxId) {
  enabledSteps(qs.indexOf(boxId) + 1, true);
}

function closeBox(boxId) {
  choices[boxId] = null;

  enabledSteps(qs.indexOf(boxId) + 1, false);
}

function displayLicences() {
  loadedLicenceData.forEach(calculateScoresForLicence);
  scores.sort(sortScores);
  console.log("displayLicences scores", scores);

  const score_list = {};

  for (var i = 0; i < scores.length; i++)
    score_list[scores[i].title] = calculateScore(
      loadedLicenceData.find(function (item) {
        return item.title === scores[i].title;
      })
    ).text;
}

function calculateScoresForLicence(licenceData) {
  console.log("calculateScoresForLicence , licenceData", licenceData);
  let nrAnswers = 0,
    nrMatches = 0,
    score = -1;
  console.log("qs", qs, choices);
  qs.forEach(function (item) {
    console.log("item", item);
    let fullChoice = choices[item];
    console.log("fullChoice", fullChoice);
    console.log(fullChoice == null);
    if (fullChoice == null) return;

    var myChoice = fullChoice.split("_")[0];

    // choice made

    console.log("myChoice", myChoice);
    if (myChoice != -1) {
      nrAnswers++;
      console.log("myChoice != -1 licenceData", licenceData)
      nrMatches = calculateQuestion(fullChoice, licenceData, nrMatches);
    }
  });

  if (nrAnswers > 0) score = nrMatches / nrAnswers;

  console.log("nrAnswers", nrAnswers, score, nrMatches);
  scores.forEach(function (item) {
    if (item.title === licenceData.title) item.score = score;
  });
  console.log("scores", scores);
}

function calculateQuestion(fullChoice, licenceData, nrMatches) {
  fullChoice = fullChoice.split("_");
  console.log("calculateQuestion fullChoice", fullChoice,licenceData);
  if (simpleYesNoQs.includes(fullChoice[0]))
    nrMatches += processSimpleYesNo(fullChoice[0], fullChoice[1], licenceData);
  else if (isLimitingQuestion(fullChoice.join("_")))
    nrMatches += processLimitingQuestion(fullChoice, licenceData);
  else if (conditionsReuseQs.includes(fullChoice[0]))
    nrMatches += processConditionsOnReuseQuestion(
      fullChoice[0],
      fullChoice[1],
      licenceData
    );

  return nrMatches;
}

function processSimpleYesNo(simpleQid, choice, licenceData) {
  var newMatch = 0,
    licenceYes = licenceData.content.includes(simpleQid);

  if (
    (choice == 1 && licenceYes) ||
    (choice == 0 && !licenceYes) ||
    choice == DONT_CARE
  )
    newMatch++;
  console.log("processSimpleYesNo newMatch", newMatch);
  return newMatch;
}

function processConditionsOnReuseQuestion(simpleQid, choice, licenceData) {
  console.log("processConditionsOnReuseQuestion licenceData", licenceData)
  let newMatch = 0,
    questionMatch = licenceData.content.includes(simpleQid);

  if (choice == 1 && questionMatch) newMatch++;

  // set q2b and q2c to 'not applicable'
  if (simpleQid == TWO_NO_COPYLEFT && choice == 0 && !questionMatch) newMatch++;
  console.log("processConditionsOnReuseQuestion newMatch", newMatch);
  return newMatch;
}

function sortScores(a, b) {
  return a.score < b.score
    ? 1
    : a.score > b.score
    ? -1
    : a.title < b.title
    ? -1
    : a.title > b.title
    ? 1
    : 0;
}


function calculateScore(licenceData) {

  var scoreText = 0, nrAnswers = 0, nrMatches = 0;

  qs.forEach(function (item) {
    var fullChoice = choices[item];

    if (!(fullChoice != null)) return;

    var myChoice = fullChoice.split('_')[0];

    if (![-1, 'na'].includes(myChoice)) {
      // choice made
      nrAnswers++;
      nrMatches = calculateQuestion(fullChoice, licenceData, nrMatches);
    }
  });

  if (nrAnswers == 0) {
    //scoreText += "No score";
    scoreText += 0;
  } else {
    //scoreText += "<span class= \"nummatches\">" + nrMatches + "</span> out of " + nrAnswers;
    scoreText += parseInt((nrMatches / nrAnswers) * 20) * 5;
  }

  return {
    matches: nrMatches,
    answers: nrAnswers,
    text: scoreText
  };
}


initLicences(stuOut);
processChoice("q1", "q1strong_1");
processChoice("q2", "q2anocopyleft_0");
processChoice("q2b", "q2bstrong_1");
//console.log("loadedLicenceData", loadedLicenceData)