var myApp = angular.module('myApp',[]);

myApp.controller('DemoController', ['$scope','$http', function($scope, $http) {
  $http.get('/view/Env').success(function(data){
  	$scope.Env=data
  	$scope.display = function(en) {
  						// alert(en.Hobbies[0]);

                   	return $scope.Hobbies=en.Hobbies;

                 }
});

  $http.get('/view/App').success(function(data){
  	$scope.App=data

  	$scope.display=function(ap) {
  		return $scope.appName=ap.appName
  	}

  });
}]);


