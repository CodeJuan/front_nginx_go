var testApp = angular.module('testApp', ['ngRoute','ngAnimate']);

testApp.config(function($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'assets/html/test.html',
            controller: 'testController'
        })
        .when('/get', {
            templateUrl: 'assets/html/get.html',
            controller: 'getController'
        })
        .when('/post', {
            templateUrl: 'assets/html/post.html',
            controller: 'postController'
        })
        .when('/bootstrap_test', {
            templateUrl: 'assets/html/bootstrap_test.html',
            controller: 'bootstrap_testController'
        });
});

testApp.controller('testController', function($scope, $http) {
    $scope.pageClass = 'test';
    $scope.routes = {}
    $scope.downData = {}
    $scope.upData = {}

    $scope.checkObject = function(obj) {
        return angular.equals({}, obj);
    };

    $scope.get_route_by_name = function() {
        $scope.routes = {}
        $scope.downData = {}
        $scope.upData = {}

        $http.get('api/v1/routes?name=' + $scope.input_name)
        .success(function(result) {
                $scope.routes = result
            }
        );
    };

    $scope.show_bus = function(routeID) {
        $scope.curID = routeID
        $scope.downData = {}
        $scope.upData = {}
            $http.get('api/v1/routeID?id=' + routeID)
            .success(function(result) {
                    $scope.route_details = result
                    $scope.downData = result.downData
                    $scope.upData = result.upData
                }
            );
        };
});

testApp.controller('getController', function($scope, $http) {
    $scope.pageClass = 'get';

        $scope.getdata = function() {
            $http.get('api/v1/users')
                .success(function(result) {
                    $scope.users = result
                }
            );
        };
});

testApp.controller('postController', function($scope, $http) {
    $scope.pageClass = 'post';

    $scope.postdata = function() {
        var data = {
                name: $scope.input_name,
                age: parseInt($scope.input_age)
            }
        $http.post('/api/v1/user/', data)
            .success(function (data, status, headers, config) {
                $scope.return_val = status
            })
            .error(function (data, status, header, config) {
                $scope.return_val = status
            });
    };
});

testApp.controller('bootstrap_testController', function($scope, $http) {
    $scope.pageClass = 'bootstrap_test'
});
