//在这一层装配好之前ltdyibiaowidget.cpp和ltddatewidget.cpp的内容
//view层与entity层的区别在于，view层的每个实体都会在前端软件中一一对应直观的显示(体现出来)
//而entity层的逻辑则不会暴露给前端的显示层
//类似如下结构:
/*
{
	widget1{
		io{

		}
		wsd{

		}
		ups{

		}
		zndb{

		}
	}
	
	widget2{
		io{

		}
		wsd{

		}
		ups{

		}
		zndb{

		}
	}
}
*/