<?php
namespace App\controllers;
use App\lib\controllers\controller;
use App\lib\entities\employee;
use App\lib\core\address;

class update extends controller
{
    public function doTask(array $params)
    {
        //0.Extract the param filter
        $emp = new employee($params[0]["empid"],$params[0]["name"],$params[0]["family"],$params[0]["phone"],$params[0]["email"],);
        //1. Fetch data from Data provider
        $res = $this->dataProvider->updateEmployee($emp);
        //2. Check the result of update operation
        if($res)
        {
            $this->redirect(address::rootaddress);
        }
        else
        {
            $this->redirect(address::rootaddress . "?task=edit&empid=" . $emp->id);
        }
    }
}

?>