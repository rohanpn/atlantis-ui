var myApp =  angular.module('myApp',[]);

myApp.controller('DemoController', ['$scope', '$http', function($scope, $http) {
	$scope.selectedApp = {};
 //    $http.get('/view/Env').success(function(data){
 //  		$scope.Env = data
 //  		$scope.display = function (en) {
 //       	return $scope.Hobbies = en.Hobbies;
 //   		}
	// });

	$http.get('/view/App').success(function(data){
		$scope.App = data;
		// $scope.selectedApp = data;
		$scope.show = function (data) {
			// load templte with particular context ap
			 $scope.selectedApp = data;
		}
	});

}]);

myApp.directive('myCustomer', function(){

	return {
		templateUrl : 'template/AppEnv.html'
	};

});

// myApp.directive('myCustomer', function() {
// 	console.log("************")

// 	var directive
// 	directive.templateUrl : 'template/AppEnv.html'
// 	return directive

// });
