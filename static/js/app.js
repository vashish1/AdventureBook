var _CONTENT=document.getElementById("heading");
//Heading of the topic
var _PART_INDEX = 0;

// Holds the handle returned from setInterval
var _INTERVAL_VAL;

// Element that holds the text
var _ELEMENT = document.querySelector("#heading");

// Implements typing effect
function Type() { 
	// Get substring with 1 characater added
	var text =  _CONTENT.substring(0, _PART_INDEX + 1);
	_ELEMENT.innerHTML = text;
	_PART_INDEX++;

	// If full sentence has been displayed then start to delete the sentence after some time
	if(text === _CONTENT) {

		clearInterval(_INTERVAL_VAL);
		setTimeout(function() {
			_INTERVAL_VAL = setInterval(Delete, 100);
		}, 1000);
	}
}

// Implements deleting effect
function Delete() {
  var text =  _CONTENT.substring(0, _PART_INDEX - 1);
	_ELEMENT.innerHTML = text;
	_PART_INDEX--;

	// If sentence has been deleted then start to display again the sentence
	if(text === '') {
		clearInterval(_INTERVAL_VAL);
		setTimeout(function() {
			_INTERVAL_VAL = setInterval(Type, 100);
		}, 200);
  }
}

// Start the typing effect on load
_INTERVAL_VAL = setInterval(Type, 100);